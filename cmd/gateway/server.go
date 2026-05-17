package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"food-delivery/internal/auth"
	"food-delivery/proto"
	"github.com/gin-gonic/gin"
)

type gatewayAuthServer struct {
	service *auth.Service
}

func (s *gatewayAuthServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.AuthResponse, error) {
	user, token, err := s.service.Register(ctx, req.Email, req.Password, req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.AuthResponse{Token: token, User: &proto.User{Id: user.ID, Email: user.Email, Name: user.Name, Role: user.Role}}, nil
}

func (s *gatewayAuthServer) Login(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	user, token, err := s.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.AuthResponse{Token: token, User: &proto.User{Id: user.ID, Email: user.Email, Name: user.Name, Role: user.Role}}, nil
}

func (s *gatewayAuthServer) GetProfile(ctx context.Context, req *proto.UserProfileRequest) (*proto.UserProfileResponse, error) {
	user, err := s.service.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &proto.UserProfileResponse{User: &proto.User{Id: user.ID, Email: user.Email, Name: user.Name, Role: user.Role}}, nil
}

func authMiddleware(s *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}
		_, claims, err := s.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		userID, ok := claims["sub"].(string)
		if !ok || userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}

func adminMiddleware(s *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authMiddleware(s)(c)
		if c.IsAborted() {
			return
		}
		_, claims, err := s.ParseToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			return
		}
		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			return
		}
		c.Next()
	}
}

func rateLimiterMiddleware() gin.HandlerFunc {
	limiter := newRateLimiter(100, 1*time.Second)
	return func(c *gin.Context) {
		if !limiter.Allow(c.ClientIP()) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		}
		c.Next()
	}
}

type rateLimiter struct {
	buckets map[string]*tokenBucket
}

type tokenBucket struct {
	tokens int
	rate   int
	last   time.Time
}

func newRateLimiter(rate int, interval time.Duration) *rateLimiter {
	return &rateLimiter{buckets: make(map[string]*tokenBucket)}
}

func (r *rateLimiter) Allow(key string) bool {
	bucket, ok := r.buckets[key]
	if !ok {
		bucket = &tokenBucket{tokens: 100, rate: 1, last: time.Now()}
		r.buckets[key] = bucket
	}
	now := time.Now()
	elapsed := now.Sub(bucket.last).Seconds()
	bucket.tokens += int(elapsed) * bucket.rate
	if bucket.tokens > 100 {
		bucket.tokens = 100
	}
	bucket.last = now
	if bucket.tokens <= 0 {
		return false
	}
	bucket.tokens--
	return true
}

func getUserID(c *gin.Context) (string, error) {
	if val, ok := c.Get("user_id"); ok {
		if id, ok := val.(string); ok {
			return id, nil
		}
	}
	return "", errors.New("user id missing")
}

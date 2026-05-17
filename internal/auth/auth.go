package auth

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type User struct {
	ID       string
	Email    string
	Name     string
	Password string
	Role     string
}

type Service struct {
	DB        *sql.DB
	JWTSecret string
}

func NewService(db *sql.DB, secret string) *Service {
	return &Service{DB: db, JWTSecret: secret}
}

func (s *Service) Register(ctx context.Context, email, password, name string) (*User, string, error) {
	hashed := hashPassword(password)
	id := uuid.NewString()
	_, err := s.DB.ExecContext(ctx, `INSERT INTO users (id, email, name, password, role) VALUES ($1, $2, $3, $4, $5)`, id, email, name, hashed, "customer")
	if err != nil {
		return nil, "", err
	}
	user := &User{ID: id, Email: email, Name: name, Role: "customer"}
	token, err := s.GenerateToken(user)
	return user, token, err
}

func (s *Service) Login(ctx context.Context, email, password string) (*User, string, error) {
	user := &User{}
	err := s.DB.QueryRowContext(ctx, `SELECT id, email, name, password, role FROM users WHERE email = $1`, email).Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Role)
	if err != nil {
		return nil, "", err
	}
	if user.Password != hashPassword(password) {
		return nil, "", errors.New("invalid credentials")
	}
	token, err := s.GenerateToken(user)
	return user, token, err
}

func (s *Service) GetUser(ctx context.Context, userID string) (*User, error) {
	user := &User{}
	err := s.DB.QueryRowContext(ctx, `SELECT id, email, name, role FROM users WHERE id = $1`, userID).Scan(&user.ID, &user.Email, &user.Name, &user.Role)
	return user, err
}

func (s *Service) GenerateToken(user *User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"name":  user.Name,
		"role":  user.Role,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.JWTSecret))
}

func (s *Service) ParseToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.JWTSecret), nil
	})
	return token, claims, err
}

func hashPassword(raw string) string {
	h := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(h[:])
}

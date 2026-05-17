package infra

import (
    "github.com/go-redis/redis/v8"
)

func NewRedisClient(redisURL string) (*redis.Client, error) {
    opt := &redis.Options{
        Addr: redisURL,
    }
    return redis.NewClient(opt), nil
}

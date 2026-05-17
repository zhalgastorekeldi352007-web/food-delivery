package infra

import (
	"context"
	"strings"

	"github.com/redis/go-redis/v9"
)

func NewRedis(ctx context.Context, raw string) (*redis.Client, error) {
	opts, err := redis.ParseURL(raw)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}

func CacheKey(parts ...string) string {
	return strings.Join(parts, ":")
}

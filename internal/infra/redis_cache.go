package infra

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{Client: client}
}

func (c *RedisCache) Set(ctx context.Context, key string, value any) error {
	buf, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Client.Set(ctx, key, buf, 0).Err()
}

func (c *RedisCache) Get(ctx context.Context, key string, dest any) error {
	data, err := c.Client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

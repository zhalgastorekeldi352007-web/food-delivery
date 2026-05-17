package infra

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestRedisCacheSetGet(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		t.Skip("redis not available")
	}
	cache := NewRedisCache(client)
	type payload struct {
		Message string `json:"message"`
	}
	value := payload{Message: "hello"}
	require.NoError(t, cache.Set(ctx, "test:cache", value))
	var got payload
	require.NoError(t, cache.Get(ctx, "test:cache", &got))
	require.Equal(t, value.Message, got.Message)
}

package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func handleRedis(c redis.Cmdable) (string, error) {
	return c.Get(context.Background(), "redis").Result()
}

func conn() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
}

func handleOrder(c redis.Cmdable) (string, error) {
	ctx := context.Background()
	_, _ = c.Get(ctx, "redis").Result()
	return c.Get(ctx, "redis").Result()
}

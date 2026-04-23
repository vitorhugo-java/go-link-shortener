package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vitorhugo-java/go-link-shortener/internal/config"
)

const redisTTL = 24 * time.Hour

func NewRedis(cfg *config.Config) *redis.Client {
	opts, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		opts = &redis.Options{Addr: "localhost:6379"}
	}
	return redis.NewClient(opts)
}

func CacheSet(rdb *redis.Client, slug, url string) error {
	return rdb.Set(context.Background(), slug, url, redisTTL).Err()
}

func CacheGet(rdb *redis.Client, slug string) (string, error) {
	return rdb.Get(context.Background(), slug).Result()
}

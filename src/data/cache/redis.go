package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           0,
		DialTimeout:  cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  cfg.Redis.PoolTimeout,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil

}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}

func Set[T any](c *redis.Client, ctx context.Context, key string, value T, duration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, v, duration).Err()
}

func Get[T any](c *redis.Client, ctx context.Context, key string) (T, error) {
	var destination T = *new(T)

	v, err := c.Get(ctx, key).Result()
	if err != nil {
		return destination, err
	}
	err = json.Unmarshal([]byte(v), destination)
	if err != nil {
		return destination, err
	}

	return destination, nil

}

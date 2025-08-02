package databases

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/BoomTHDev/golang_clean_arch/config"
	"github.com/redis/go-redis/v9"
)

var (
	onceRedis   sync.Once
	redisClient RedisClient
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(conf *config.Redis) RedisClient {
	onceRedis.Do(func() {
		opt, err := redis.ParseURL(conf.URL)
		if err != nil {
			log.Printf("Warning: Invalid Redis URL: %v", err)
		}
		client := redis.NewClient(opt)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if _, err := client.Ping(ctx).Result(); err != nil {
			panic(err)
		}
		redisClient = RedisClient{client: client}
	})
	return redisClient
}

func (r *RedisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisClient) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

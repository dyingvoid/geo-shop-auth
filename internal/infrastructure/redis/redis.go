package redis

import "github.com/redis/go-redis/v9"

func NewRedis(opts RedisOptions) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     opts.Addr,
		Password: opts.Password,
		DB:       opts.DB,
	})
	return rdb
}

type RedisOptions struct {
	Addr     string
	Password string
	DB       int
}

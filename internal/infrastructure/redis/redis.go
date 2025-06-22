package redis

import "github.com/redis/go-redis/v9"

func NewRedis(opts Options) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     opts.Addr,
		Password: opts.Password,
		DB:       opts.DB,
	})
	return rdb
}

type Options struct {
	Addr     string
	Password string
	DB       int
}

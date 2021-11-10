package redis

import (
	"github.com/go-redis/redis/v8"
)

type Option func(*redis.Options)

func SetDefault() redis.Options {
	return redis.Options{
		Addr:       "127.0.0.1:6379",
		Username:   "",
		Password:   "",
		DB:         0,
		PoolSize:   16,
		MaxRetries: 3,
	}
}

func SetAddr(addr string) Option {
	return func(o *redis.Options) {
		o.Addr = addr
	}
}

func SetUsername(username string) Option {
	return func(o *redis.Options) {
		o.Username = username
	}
}

func SetPassword(password string) Option {
	return func(o *redis.Options) {
		o.Password = password
	}
}

func SetDB(db int) Option {
	return func(o *redis.Options) {
		o.DB = db
	}
}

func SetPoolSize(poolSize int) Option {
	return func(o *redis.Options) {
		o.PoolSize = poolSize
	}
}

func SetMaxRetries(maxRetries int) Option {
	return func(o *redis.Options) {
		o.MaxRetries = maxRetries
	}
}

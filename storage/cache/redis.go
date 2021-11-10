package cache

import (
	yredis "gitlab.benlai.work/go/ymir/storage/redis"
	"time"
)

// NewRedis redis模式
func NewRedisCache(opts ...yredis.Option) (*yredis.Redis, error) {
	op := yredis.SetDefault()
	for _, o := range opts {
		o(&op)
	}

	r, err := yredis.NewWithOptions(nil, &op)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Redis cache implement
type RedisCache struct {
	yredis *yredis.Redis
}

func (*RedisCache) String() string {
	return "redis"
}

// connect connect test
func (r *RedisCache) connect() error {
	return r.yredis.Ping()
}

// Get from key
func (r *RedisCache) Get(key string) (string, error) {
	return r.yredis.Get(key)
}

// Set value with key and expire time
func (r *RedisCache) Set(key string, val interface{}, expire int) error {
	_, err := r.yredis.SetWithExpire(key, val, expire)
	return err
}

// Del delete key in redis
func (r *RedisCache) Del(key string) error {
	_, err := r.yredis.Del(key)
	return err
}

// HashGet from key
func (r *RedisCache) HashGet(hk, key string) (string, error) {
	return r.yredis.HGet(hk, key)
}

// HashDel delete key in specify redis's hashtable
func (r *RedisCache) HashDel(hk, key string) error {
	_, err := r.yredis.HDel(hk, key)
	return err
}

// Increase
func (r *RedisCache) Increase(key string) (int64, error) {
	return r.yredis.Incr(key)
}

func (r *RedisCache) Decrease(key string) (int64, error) {
	return r.yredis.Decr(key)
}

// Set ttl
func (r *RedisCache) Expire(key string, dur time.Duration) (bool, error) {
	return r.yredis.Expire(key, dur)
}

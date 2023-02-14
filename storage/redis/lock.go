package redis

import (
	"github.com/bsm/redislock"
	"time"
)

func (r *Redis) Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error) {
	if r.mutex == nil {
		r.mutex = redislock.New(r.client)
	}
	return r.mutex.Obtain(r.ctx, key, time.Duration(ttl)*time.Second, options)
}

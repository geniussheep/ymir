package cache

import "time"

type AdapterCache interface {
	String() string
	Get(key string) (string, error)
	Set(key string, val interface{}, expire int) error
	Del(key string) error
	HashGet(hk, key string) (string, error)
	HashDel(hk, key string) error
	Increase(key string) (int64, error)
	Decrease(key string) (int64, error)
	Expire(key string, dur time.Duration) (bool, error)
}

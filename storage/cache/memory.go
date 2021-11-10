package cache

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cast"
)

type item struct {
	Value   string
	Expired time.Time
}

// NewMemory memory模式
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		items: new(sync.Map),
	}
}

type MemoryCache struct {
	items *sync.Map
	mutex sync.RWMutex
}

func (*MemoryCache) String() string {
	return "memory_cache"
}

func (m *MemoryCache) connect() {
}

func (m *MemoryCache) Get(key string) (string, error) {
	item, err := m.getItem(key)
	if err != nil || item == nil {
		return "", err
	}
	return item.Value, nil
}

func (m *MemoryCache) getItem(key string) (*item, error) {
	var err error
	i, ok := m.items.Load(key)
	if !ok {
		return nil, nil
	}
	switch i.(type) {
	case *item:
		item := i.(*item)
		if item.Expired.Before(time.Now()) {
			//过期
			_ = m.del(key)
			//过期后删除
			return nil, nil
		}
		return item, nil
	default:
		err = fmt.Errorf("value of %s type error", key)
		return nil, err
	}
}

func (m *MemoryCache) Set(key string, val interface{}, expire int) error {
	s, err := cast.ToStringE(val)
	if err != nil {
		return err
	}
	item := &item{
		Value:   s,
		Expired: time.Now().Add(time.Duration(expire) * time.Second),
	}
	return m.setItem(key, item)
}

func (m *MemoryCache) setItem(key string, item *item) error {
	m.items.Store(key, item)
	return nil
}

func (m *MemoryCache) Del(key string) error {
	return m.del(key)
}

func (m *MemoryCache) del(key string) error {
	m.items.Delete(key)
	return nil
}

func (m *MemoryCache) HashGet(hk, key string) (string, error) {
	item, err := m.getItem(hk + key)
	if err != nil || item == nil {
		return "", err
	}
	return item.Value, err
}

func (m *MemoryCache) HashDel(hk, key string) error {
	return m.del(hk + key)
}

func (m *MemoryCache) Increase(key string) (int64, error) {
	return m.calculate(key, 1)
}

func (m *MemoryCache) Decrease(key string) (int64, error) {
	return m.calculate(key, -1)
}

func (m *MemoryCache) calculate(key string, num int64) (int64, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	item, err := m.getItem(key)
	if err != nil {
		return 0, err
	}

	if item == nil {
		err = fmt.Errorf("%s not exist", key)
		return 0, err
	}
	var n int64
	n, err = cast.ToInt64E(item.Value)
	if err != nil {
		return 0, err
	}
	n += num
	item.Value = strconv.Itoa(int(n))
	return n, m.setItem(key, item)
}

func (m *MemoryCache) Expire(key string, dur time.Duration) (bool, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	item, err := m.getItem(key)
	if err != nil {
		return false, err
	}
	if item == nil {
		err = fmt.Errorf("%s not exist", key)
		return false, err
	}
	item.Expired = time.Now().Add(dur)
	return true, m.setItem(key, item)
}

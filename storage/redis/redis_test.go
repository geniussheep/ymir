package redis

import (
	"github.com/go-redis/redis/v8"
	"gitlab.benlai.work/go/ymir/logger"
	"testing"
)

func TestDb(t *testing.T) {
	r, err := New(nil, &redis.Options{
		Addr: "192.168.60.12:6379",
		DB:   0,
	})
	if err != nil {
		logger.Errorf("new redis error:%e", err.Error())
	}

	rst, err := r.Set("test-yy", "123")
	if err != nil {
		logger.Errorf("redis set error:%e", err.Error())
	}
	logger.Infof("set result:%s", rst)
	rsg, err := r.Get("test-yy")
	if err != nil {
		logger.Errorf("redis get error:%e", err.Error())
	}

	logger.Infof("get test-yy:%s", rsg)
}

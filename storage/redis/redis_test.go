package redis

import (
	"gitlab.benlai.work/go/ymir/logger"
	"testing"
)

func TestDb(t *testing.T) {
	r, err := New(
		SetAddr("10.250.60.11:6379"),
		SetDB(0),
	)
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

package service

import (
	"fmt"
	"gitlab.benlai.work/go/ymir/component/zookeeper"
	"gitlab.benlai.work/go/ymir/storage/redis"

	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/storage/db"
)

type Service struct {
	Orm       map[string]*db.Yorm
	Redis     map[string]*redis.Redis
	Zookeeper map[string]*zookeeper.Zookeeper
	Log       *logger.Helper
	Error     error
}

func (svc *Service) AddError(err error) error {
	if svc.Error == nil {
		svc.Error = err
	} else if err != nil {
		svc.Error = fmt.Errorf("%v; %w", svc.Error, err)
	}
	return svc.Error
}

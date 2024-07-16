package service

import (
	"fmt"
	"github.com/geniussheep/ymir/component/zookeeper"
	"github.com/geniussheep/ymir/k8s"
	"github.com/geniussheep/ymir/logger"
	"github.com/geniussheep/ymir/storage/db"
	"github.com/geniussheep/ymir/storage/redis"
)

type Service struct {
	Orm       map[string]*db.Yorm
	Redis     map[string]*redis.Redis
	Zookeeper map[string]*zookeeper.Zookeeper
	Other     map[string]interface{}
	K8S       map[string]*k8s.Client
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

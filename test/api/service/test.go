package service

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.benlai.work/go/ymir/sdk/service"
	"gitlab.benlai.work/go/ymir/test/api/model"
	"gitlab.benlai.work/go/ymir/test/api/service/dto"
	"gitlab.benlai.work/go/ymir/test/common"
)

type Test struct {
	service.Service
}

func (s *Test) GetApplication(q *dto.QueryApp) (*model.Application, error) {
	if err := q.CheckArgs(); err != nil {
		return nil, err
	}
	var model model.Application
	err := s.Orm[common.DbMonitor].FindOne(cast.ToInt64(q.AppId), &model)
	if err != nil {
		s.Log.Errorf("get app by appId:%s error:%s", q.AppId, err)
		return nil, err
	}
	if !model.Validate() {
		return nil, fmt.Errorf("there is no app[id: %d] info in db", q.AppId)
	}
	rd, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}
	s.Redis[common.TestRedis].Set("TestApp", string(rd))
	err = s.Zookeeper[common.GetZkName(q.Environment)].Notify("/application/base", nil)
	if err != nil {
		return nil, err
	}

	testModel, err := s.Redis[common.TestRedis].Get("TestApp")
	if err != nil {
		return nil, err
	}
	s.Log.Info("test application", testModel)
	return &model, nil
}

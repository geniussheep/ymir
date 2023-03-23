package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab.benlai.work/go/ymir/sdk/api"
	"gitlab.benlai.work/go/ymir/test/api/model"
	"gitlab.benlai.work/go/ymir/test/api/service"
	"gitlab.benlai.work/go/ymir/test/api/service/dto"
	"gitlab.benlai.work/go/ymir/test/common"
)

type Test struct {
	api.Api
}

// Get
// @Summary 测试Controller
// @Description 测试Controller
// @Tags 测试Controller
// @Param pathArgs path string false "url内参数"
// @Param qString query string false "字符串, 如:string"
// @Param qInt body int false "数字 如:123"
// @Success 200 {object} response.Response{data=[]model.RespTest} "{"code": 200, "data": [...]}"
// @Router /api/v1/test/{pathArgs}/get [get]
// @Security Bearer
func (api Test) Get(c *gin.Context) {
	req := model.QueryTest{}
	err := api.MakeContext(c).
		Bind(&req).
		Bind(&req, binding.Query, binding.JSON).
		Errors
	if err != nil {
		api.Error(500, err, err.Error())
		return
	}

	resp := model.RespTest{QueryTest: req, RespBody: "yangyangbody", RespType: "yangyang--respType"}
	api.OK(resp, "test get")
}

// GetApplication
// @Summary 获取App信息
// @Description 获取App信息
// @Tags 应用信息管理
// @Param appId path int false "应用Id"
// @Param environment path string false "应用环境"
// @Success 200 {object} response.Response{data=[]model.Application} "{"code": 200, "data": [...]}"
// @Router /api/v1/app/get/{appId}/{environment} [get]
// @Security Bearer
func (api Test) GetApplication(c *gin.Context) {
	req := dto.QueryApp{}
	s := service.Test{}
	err := api.MakeContext(c).
		Bind(&req).
		MakeOrm(common.DbMonitor).
		MakeRedis(common.TestRedis).
		MakeZookeeper(common.GetZkName(req.Environment)).
		MakeService(&s.Service).
		Errors

	if err != nil {
		api.Error(500, err, err.Error())
		return
	}
	if err := req.CheckArgs(); err != nil {
		api.Error(500, err, err.Error())
		return
	}
	objects, err := s.GetApplication(&req)
	if err != nil {
		api.Error(500, err, err.Error())
		return
	}
	api.OK(objects, "get application success")
}

// CreateByAppId 创建Jenkins编译任务
// @Summary 创建Jenkins编译任务
// @Description 创建Jenkins编译任务
// @Tags Jenkins应用编译任务管理
// @Param appId path string true "应用Id"
// @Param deployType formData string true "任务类型 [Normal:完整发布, BuildImage:制作镜像, GoLive:应用上线, RestartApp:重启发布, Config:配置发布, ConfigRestart:配置发布, ConfigWithoutRestart:配置发布, QuicklyDeploy:快速部署, QuicklyRollback:快速回滚, Rollback:Rollback]"
// @Param environment formData string true "环境"
// @Param isAutoPublish formData bool false "是否全自动发布"
// @Success 200 {object} response.Response{data=string} "{"code": 200, "data": [...]}"
// @Router /api/v1/app/id/{appId}/build/create [post]
// @Security Bearer
func (api Test) CreateByAppId(c *gin.Context) {
	req := dto.CreateBuild{}
	s := service.Test{}
	err := api.MakeContext(c).
		Bind(&req).
		Bind(&req, binding.Query).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		api.Logger.Error(err)
		api.Error(500, err, err.Error())
		return
	}

	api.OK(req, fmt.Sprintf("test create app build job by appId:%s success", req.AppId))
}

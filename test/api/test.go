package api

import (
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
// @Param queryString formData string false "字符串, 如:string"
// @Param queryInt formData int false "数字 如:123"
// @Success 200 {object} response.Response{data=[]model.RespTest} "{"code": 200, "data": [...]}"
// @Router /api/v1/test/{pathArgs}/get [get]
// @Security Bearer
func (api Test) Get(c *gin.Context) {
	req := model.QueryTest{}
	err := api.MakeContext(c).
		MakeOrm("benlaimonitor").
		Bind(&req).
		Bind(&req, binding.Query, binding.Form, binding.JSON).
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
// @Success 200 {object} response.Response{data=[]model.Application} "{"code": 200, "data": [...]}"
// @Router /api/v1/app/get/{appId} [get]
// @Security Bearer
func (api Test) GetApplication(c *gin.Context) {
	req := dto.QueryByAppId{}
	s := service.Test{}
	err := api.MakeContext(c).
		MakeOrm(common.DbMonitor).
		MakeRedis(common.TestRedis).
		Bind(&req).
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

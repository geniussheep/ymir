package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab.benlai.work/go/ymir/cli"
	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/sdk"
	"gitlab.benlai.work/go/ymir/sdk/api"
	"gitlab.benlai.work/go/ymir/sdk/cmd"
	"gitlab.benlai.work/go/ymir/sdk/common"
	"gitlab.benlai.work/go/ymir/sdk/middleware"
	"log"
	"os"
)

const (
	V1 = "/api/v1"
)

var (
	TestApi       = &api.Api{}
	ApiRouterScan = make([]func(), 0)
)

func init() {
	ApiRouterScan = append(ApiRouterScan, func() {
		app := TestController{}
		TestApi.AppendRouters(V1, api.RouterEntry{Path: "test/{path}/get ", Method: "GET", Handler: app.Get})
	})
	for _, routerScan := range ApiRouterScan {
		routerScan()
	}
}

type TestController struct {
	api.Api
}

type QueryTest struct {
	Path    string `msgpack:"path" json:"path" form:"path" xml:"path" uri:"path" example:"path"`
	QString string `msgpack:"qString" json:"qString" form:"qString" xml:"qString" uri:"qString" example:"string"`
	QInt    int    `msgpack:"qInt" json:"qInt" form:"qInt" xml:"qInt" uri:"qInt"  binding:"required" example:"123"`
}

// Get
// @Summary 测试Controller
// @Description 测试Controller
// @Tags 测试Controller
// @Param path path string false "url内参数"
// @Param queryString query string false "字符串, 如:string"
// @Param queryInt query int false "数字 如:123"
// @Success 200 {object} response.Response{data=[]model.QueryTest} "{"code": 200, "data": [...]}"
// @Router /api/v1/test/{path}/get [get]
// @Security Bearer
func (api TestController) Get(c *gin.Context) {
	req := QueryTest{}
	err := api.MakeContext(c).
		Bind(&req).
		Bind(&req, binding.Query).Errors
	if err != nil {
		api.Error(500, err, err.Error())
		return
	}
	api.OK(req, "test get")

}

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		logger.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	middleware.InitMiddleware(r)

	TestApi.RegisterRouters(r)
}

func main() {
	p := &cli.Program{
		Program:        "TestCmd",
		Desc:           "test cmd",
		AppRoutersScan: []func(){InitRouter},
		ConfigFilePath: common.DEFAULT_CONFIG_FILE_PATH,
		ExtendConfig:   nil,
	}
	rootCmd := cmd.New(p)
	rootCmd.Execute()
}

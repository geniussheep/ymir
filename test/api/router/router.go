package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/sdk"
	"gitlab.benlai.work/go/ymir/sdk/api"
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

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	for _, routerScan := range ApiRouterScan {
		routerScan()
	}
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

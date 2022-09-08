package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab.benlai.work/go/ymir/sdk/api/response"
	"gitlab.benlai.work/go/ymir/sdk/config"
	"gitlab.benlai.work/go/ymir/sdk/pkg"
	"gitlab.benlai.work/go/ymir/sdk/service"
	"net/http"

	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/storage/db"
)

type Api struct {
	Context *gin.Context
	Logger  *logger.Helper
	Orm     map[string]*db.Yorm
	Routers map[string][]RouterEntry
	Errors  error
}

func (api *Api) AddError(err error) {
	if api.Errors == nil {
		api.Errors = err
	} else if err != nil {
		api.Logger.Error(err)
		api.Errors = fmt.Errorf("%v; %w", api.Errors, err)
	}
}

// MakeContext 设置http上下文
func (api *Api) MakeContext(c *gin.Context) *Api {
	api.Context = c
	api.Logger = GetRequestLogger(c)
	return api
}

// GetLogger 获取上下文提供的日志
func (api Api) GetLogger() *logger.Helper {
	return GetRequestLogger(api.Context)
}

func (api *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		err = api.Context.ShouldBindUri(d)
		if err != nil {
			api.AddError(err)
		}
	} else {
		for i := range bindings {
			switch bindings[i] {
			case binding.JSON:
				err = api.Context.ShouldBindWith(d, binding.JSON)
			case binding.XML:
				err = api.Context.ShouldBindWith(d, binding.XML)
			case binding.Form:
				err = api.Context.ShouldBindWith(d, binding.Form)
			case binding.Query:
				err = api.Context.ShouldBindWith(d, binding.Query)
			case binding.FormPost:
				err = api.Context.ShouldBindWith(d, binding.FormPost)
			case binding.FormMultipart:
				err = api.Context.ShouldBindWith(d, binding.FormMultipart)
			case binding.ProtoBuf:
				err = api.Context.ShouldBindWith(d, binding.ProtoBuf)
			case binding.MsgPack:
				err = api.Context.ShouldBindWith(d, binding.MsgPack)
			case binding.YAML:
				err = api.Context.ShouldBindWith(d, binding.YAML)
			case binding.Header:
				err = api.Context.ShouldBindWith(d, binding.Header)
			default:
				api.AddError(fmt.Errorf("api bind error: the binding type:%s unknown", bindings[i].Name()))
			}
			if err != nil {
				api.AddError(err)
			}
		}
	}
	return api
}

func (api *Api) MakeOrm(dbName string) *Api {
	if api.Orm == nil {
		api.Orm = make(map[string]*db.Yorm)
	}
	if _, ok := api.Orm[dbName]; ok {
		return api
	}
	yorm, err := GetOrm(api.Context, dbName)
	if err != nil {
		api.AddError(fmt.Errorf("set orm:[name: %s] error: %s", dbName, err))
	}
	api.Orm[dbName] = yorm
	return api
}

func (api *Api) MakeService(c *service.Service) *Api {
	c.Log = api.Logger
	c.Orm = api.Orm
	return api
}

func (api *Api) SetRouterGroup(routerPrefix string) *Api {
	if api.Routers == nil {
		api.Routers = map[string][]RouterEntry{}
	}
	if _, ok := api.Routers[routerPrefix]; !ok {
		api.Routers[routerPrefix] = make([]RouterEntry, 0)
	}
	return api
}

func (api *Api) AppendRouters(routerPrefix string, routers ...RouterEntry) *Api {
	if api.Routers == nil {
		api.SetRouterGroup(routerPrefix)
	}
	api.Routers[routerPrefix] = append(api.Routers[routerPrefix], routers...)
	return api
}

// RegisterRouters registers APIs.
func (api *Api) RegisterRouters(engine *gin.Engine) {
	isDebug := config.ApplicationConfig.Mode == pkg.Dev.String() || config.ApplicationConfig.Mode == pkg.Test.String()
	engine.Use(LoggerMiddleware(logger.NewHelper(logger.DefaultLogger), isDebug, "/scanv.htm"))

	if api.Routers == nil {
		api.AddError(fmt.Errorf("register api routers error: api.Routers is nil"))
		return
	}
	for routerPrefix, routers := range api.Routers {
		rg := engine.Group(routerPrefix)
		for _, r := range routers {
			switch r.Method {
			case "GET":
				rg.GET(r.Path, r.Handler)
			case "HEAD":
				rg.HEAD(r.Path, r.Handler)
			case "PUT":
				rg.PUT(r.Path, r.Handler)
			case "POST":
				rg.POST(r.Path, r.Handler)
			case "PATCH":
				rg.PATCH(r.Path, r.Handler)
			case "DELETE":
				rg.DELETE(r.Path, r.Handler)
			case "OPTIONS":
				rg.OPTIONS(r.Path, r.Handler)
			}
		}
	}
	engine.GET("/scanv.htm", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
}

// Error 通常错误数据处理
func (api Api) Error(code int, err error, msg string) {
	response.Error(api.Context, code, err, msg)
}

// OK 通常成功数据处理
func (api Api) OK(data interface{}, msg string) {
	response.OK(api.Context, data, msg)
}

// PageOK 分页数据处理
func (api Api) PageOK(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	response.PageOK(api.Context, result, count, pageIndex, pageSize, msg)
}

// Custom 兼容函数
func (api Api) Custom(data gin.H) {
	response.Custum(api.Context, data)
}

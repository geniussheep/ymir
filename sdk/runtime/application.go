package runtime

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/sdk/api"
	"gitlab.benlai.work/go/ymir/storage/cache"
	"gitlab.benlai.work/go/ymir/storage/db"
	"gitlab.benlai.work/go/ymir/storage/redis"
	"log"
	"net/http"
	"os"
	"sync"
)

type Application struct {
	webApi      *api.Api
	dbs         map[string]*db.Yorm
	cache       cache.AdapterCache
	redis       map[string]*redis.Redis
	casbins     map[string]*casbin.SyncedEnforcer
	middlewares map[string]interface{}
	handler     map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)
	engine      http.Handler
	mux         sync.RWMutex
	other       map[string]interface{}
}

// New 默认值
func New() *Application {
	return &Application{
		webApi:      &api.Api{},
		dbs:         make(map[string]*db.Yorm),
		cache:       cache.NewMemoryCache(),
		redis:       make(map[string]*redis.Redis),
		casbins:     make(map[string]*casbin.SyncedEnforcer),
		middlewares: make(map[string]interface{}),
		handler:     make(map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)),
	}
}

// SetWebApi 设置webapi对象
func (a *Application) SetWebApi(webapi *api.Api) {
	a.webApi = webapi
}

// GetWebApi 获取webapi对象
func (a *Application) GetWebApi() *api.Api {
	return a.webApi
}

// SetEngine 设置路由引擎
func (a *Application) SetEngine(engine http.Handler) {
	a.engine = engine
}

// GetEngine 获取路由引擎
func (a *Application) GetEngine() http.Handler {
	return a.engine
}

// GetGinEngine 获取GinEngine路由引擎
func (a *Application) GetGinEngine() *gin.Engine {
	var r *gin.Engine
	h := a.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}
	return r
}

// SetLogger 设置日志组件
func (a *Application) SetLogger(l logger.Logger) {
	logger.DefaultLogger = l
}

// GetLogger 获取日志组件
func (a *Application) GetLogger() logger.Logger {
	return logger.DefaultLogger
}

func (a *Application) SetDb(dbName string, db *db.Yorm) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.dbs[dbName] = db
}

// GetDb 获取所有map里的db数据
func (a *Application) GetDb(dbName string) *db.Yorm {
	a.mux.Lock()
	defer a.mux.Unlock()
	if _, ok := a.dbs[dbName]; !ok {
		return nil
	}
	return a.dbs[dbName]
}

func (a *Application) SetCasbin(key string, enforcer *casbin.SyncedEnforcer) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.casbins[key] = enforcer
}

func (a *Application) GetAllCasbin() map[string]*casbin.SyncedEnforcer {
	return a.casbins
}

// GetCasbin 根据key获取casbin
func (a *Application) GetCasbin(key string) *casbin.SyncedEnforcer {
	a.mux.Lock()
	defer a.mux.Unlock()
	if e, ok := a.casbins["*"]; ok {
		return e
	}
	return a.casbins[key]
}

func (a *Application) SetRedis(rName string, redis *redis.Redis) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.redis[rName] = redis
}

func (a *Application) GetRedis(rName string) *redis.Redis {
	a.mux.Lock()
	defer a.mux.Unlock()
	if _, ok := a.redis[rName]; !ok {
		return nil
	}
	return a.redis[rName]
}

// SetMiddleware 设置中间件
func (a *Application) SetMiddleware(key string, middleware interface{}) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.middlewares[key] = middleware
}

// GetAllMiddleware 获取所有中间件
func (a *Application) GetAllMiddleware() map[string]interface{} {
	return a.middlewares
}

// GetMiddleware 获取对应key的中间件
func (a *Application) GetMiddleware(key string) interface{} {
	a.mux.Lock()
	defer a.mux.Unlock()
	return a.middlewares[key]
}

// SetOtherComponent 设置其他组件实例
func (a *Application) SetOtherComponent(key string, other interface{}) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.other[key] = other
}

// GetAllOtherComponent 设置所有其他组件实例
func (a *Application) GetAllOtherComponent() map[string]interface{} {
	return a.other
}

// GetOtherComponent 获取对应key的组件
func (a *Application) GetOtherComponent(key string) interface{} {
	a.mux.Lock()
	defer a.mux.Unlock()
	return a.other[key]
}

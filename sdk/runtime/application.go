package runtime

import (
	"github.com/casbin/casbin/v2"
	"github.com/geniussheep/ymir/component/zookeeper"
	"github.com/geniussheep/ymir/k8s"
	"github.com/geniussheep/ymir/logger"
	"github.com/geniussheep/ymir/sdk/api"
	"github.com/geniussheep/ymir/storage/cache"
	"github.com/geniussheep/ymir/storage/db"
	"github.com/geniussheep/ymir/storage/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sync"
)

type Application struct {
	webApi     *api.Api
	db         map[string]*db.Yorm
	cache      cache.AdapterCache
	redis      map[string]*redis.Redis
	zookeeper  map[string]*zookeeper.Zookeeper
	casbin     map[string]*casbin.SyncedEnforcer
	middleware map[string]interface{}
	handler    map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)
	engine     http.Handler
	mux        sync.RWMutex
	other      map[string]interface{}
	k8s        map[string]*k8s.Client
}

// New 默认值
func New() *Application {
	return &Application{
		webApi:     &api.Api{},
		db:         make(map[string]*db.Yorm),
		cache:      cache.NewMemoryCache(),
		redis:      make(map[string]*redis.Redis),
		zookeeper:  make(map[string]*zookeeper.Zookeeper),
		casbin:     make(map[string]*casbin.SyncedEnforcer),
		middleware: make(map[string]interface{}),
		handler:    make(map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)),
		other:      make(map[string]interface{}),
		k8s:        make(map[string]*k8s.Client),
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
	a.db[dbName] = db
}

// GetDb 获取所有map里的db数据
func (a *Application) GetDb(dbName string) *db.Yorm {
	a.mux.Lock()
	defer a.mux.Unlock()
	if _, ok := a.db[dbName]; !ok {
		return nil
	}
	return a.db[dbName]
}

func (a *Application) SetCasbin(key string, enforcer *casbin.SyncedEnforcer) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.casbin[key] = enforcer
}

func (a *Application) GetAllCasbin() map[string]*casbin.SyncedEnforcer {
	return a.casbin
}

// GetCasbin 根据key获取casbin
func (a *Application) GetCasbin(key string) *casbin.SyncedEnforcer {
	a.mux.Lock()
	defer a.mux.Unlock()
	if e, ok := a.casbin["*"]; ok {
		return e
	}
	return a.casbin[key]
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

func (a *Application) SetZookeeper(zkName string, zk *zookeeper.Zookeeper) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.zookeeper[zkName] = zk
}

// GetZookeeper 获取所有map里的Zk数据
func (a *Application) GetZookeeper(zkName string) *zookeeper.Zookeeper {
	a.mux.Lock()
	defer a.mux.Unlock()
	if _, ok := a.zookeeper[zkName]; !ok {
		return nil
	}
	return a.zookeeper[zkName]
}

// SetMiddleware 设置中间件
func (a *Application) SetMiddleware(key string, middleware interface{}) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.middleware[key] = middleware
}

// GetAllMiddleware 获取所有中间件
func (a *Application) GetAllMiddleware() map[string]interface{} {
	return a.middleware
}

// GetMiddleware 获取对应key的中间件
func (a *Application) GetMiddleware(key string) interface{} {
	a.mux.Lock()
	defer a.mux.Unlock()
	return a.middleware[key]
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

func (a *Application) SetK8S(k8sName string, k8s *k8s.Client) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.k8s[k8sName] = k8s
}

func (a *Application) GetK8S(k8sName string) *k8s.Client {
	a.mux.Lock()
	defer a.mux.Unlock()
	if _, ok := a.k8s[k8sName]; !ok {
		return nil
	}
	return a.k8s[k8sName]
}

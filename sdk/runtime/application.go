package runtime

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/storage/cache"
	"gitlab.benlai.work/go/ymir/storage/db"
	"gitlab.benlai.work/go/ymir/storage/redis"
	"net/http"
	"sync"
)

type Application struct {
	dbs         map[string]*db.Yorm
	cache       cache.AdapterCache
	redis       map[string]*redis.Redis
	casbins     map[string]*casbin.SyncedEnforcer
	middlewares map[string]interface{}
	handler     map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)
	engine      http.Handler
	mux         sync.RWMutex
}

// NewConfig 默认值
func New() *Application {
	return &Application{
		dbs:         make(map[string]*db.Yorm),
		cache:       cache.NewMemoryCache(),
		casbins:     make(map[string]*casbin.SyncedEnforcer),
		middlewares: make(map[string]interface{}),
		handler:     make(map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)),
	}
}

// SetEngine 设置路由引擎
func (a *Application) SetEngine(engine http.Handler) {
	a.engine = engine
}

// GetEngine 获取路由引擎
func (a *Application) GetEngine() http.Handler {
	return a.engine
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

func (a *Application) GetCasbin() map[string]*casbin.SyncedEnforcer {
	return a.casbins
}

// GetCasbinKey 根据key获取casbin
func (a *Application) GetCasbinKey(key string) *casbin.SyncedEnforcer {
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
	return a.redis[rName]
}

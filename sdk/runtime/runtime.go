package runtime

import (
	"github.com/casbin/casbin/v2"
	"github.com/geniussheep/ymir/component/zookeeper"
	"github.com/geniussheep/ymir/k8s"
	"github.com/geniussheep/ymir/logger"
	"github.com/geniussheep/ymir/sdk/api"
	"github.com/geniussheep/ymir/storage/db"
	"github.com/geniussheep/ymir/storage/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Runtime interface {
	SetWebApi(webapi *api.Api)
	GetWebApi() *api.Api

	SetEngine(engine http.Handler)
	GetEngine() http.Handler
	GetGinEngine() *gin.Engine

	// SetLogger 使用go-admin定义的logger，参考来源go-micro
	SetLogger(logger logger.Logger)
	GetLogger() logger.Logger

	SetDb(dbName string, db *db.Yorm)
	GetDb(dbName string) *db.Yorm

	SetCasbin(key string, enforcer *casbin.SyncedEnforcer)
	GetAllCasbin() map[string]*casbin.SyncedEnforcer
	GetCasbin(key string) *casbin.SyncedEnforcer

	// redis
	SetRedis(rName string, redis *redis.Redis)
	GetRedis(rName string) *redis.Redis

	// zookeeper
	SetZookeeper(zkName string, zk *zookeeper.Zookeeper)
	GetZookeeper(zkName string) *zookeeper.Zookeeper

	// SetMiddleware middleware
	SetMiddleware(string, interface{})
	GetAllMiddleware() map[string]interface{}
	GetMiddleware(key string) interface{}

	SetOtherComponent(string, interface{})
	GetAllOtherComponent() map[string]interface{}
	GetOtherComponent(key string) interface{}

	SetK8S(k8sName string, client *k8s.Client)
	GetK8S(k8sName string) *k8s.Client
}

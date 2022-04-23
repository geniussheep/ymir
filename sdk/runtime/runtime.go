package runtime

import (
	"github.com/casbin/casbin/v2"
	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/storage/db"
	"gitlab.benlai.work/go/ymir/storage/redis"
	"net/http"
)

type Runtime interface {
	SetEngine(engine http.Handler)
	GetEngine() http.Handler

	// SetLogger 使用go-admin定义的logger，参考来源go-micro
	SetLogger(logger logger.Logger)
	GetLogger() logger.Logger

	SetDb(dbName string, db *db.Yorm)
	GetDb(dbName string) *db.Yorm

	SetCasbin(key string, enforcer *casbin.SyncedEnforcer)
	GetCasbin() map[string]*casbin.SyncedEnforcer
	GetCasbinKey(key string) *casbin.SyncedEnforcer

	// redis
	SetRedis(rName string, redis *redis.Redis)
	GetRedis(rName string) *redis.Redis
}

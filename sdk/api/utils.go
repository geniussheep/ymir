package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.benlai.work/go/ymir/storage/db"
	"gitlab.benlai.work/go/ymir/storage/redis"
	"strings"
)

// GetOrm 获取orm连接
func GetOrm(ctx *gin.Context, dbName string) (*db.Yorm, error) {
	idb, exist := ctx.Get(strings.ToLower(dbName))
	if !exist {
		return nil, fmt.Errorf("the db:%s connect not exist", dbName)
	}
	switch idb.(type) {
	case *db.Yorm:
		//新增操作
		return idb.(*db.Yorm).WithContext(ctx), nil
	default:
		return nil, fmt.Errorf("the db:%s connect not exist", dbName)
	}
}

// GetRedis 获取redis连接
func GetRedis(ctx *gin.Context, redisName string) (*redis.Redis, error) {
	_redis, exist := ctx.Get(strings.ToLower(redisName))
	if !exist {
		return nil, fmt.Errorf("the redis:%s connect not exist", redisName)
	}
	switch _redis.(type) {
	case *redis.Redis:
		//新增操作
		return _redis.(*redis.Redis).WithContext(ctx), nil
	default:
		return nil, fmt.Errorf("the redis:%s connect not exist", redisName)
	}
}

package api

import (
	"fmt"
	"github.com/geniussheep/ymir/component/zookeeper"
	"github.com/geniussheep/ymir/k8s"
	"github.com/geniussheep/ymir/storage/db"
	"github.com/geniussheep/ymir/storage/redis"
	"github.com/gin-gonic/gin"
)

// GetOrm 获取orm连接
func GetOrm(ctx *gin.Context, dbName string) (*db.Yorm, error) {
	idb, exist := ctx.Get(dbName)
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
	_redis, exist := ctx.Get(redisName)
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

// GetZookeeper 获取zookeeper连接
func GetZookeeper(ctx *gin.Context, zkName string) (*zookeeper.Zookeeper, error) {
	_zk, exist := ctx.Get(zkName)
	if !exist {
		return nil, fmt.Errorf("the zk:%s connect not exist", zkName)
	}
	switch _zk.(type) {
	case *zookeeper.Zookeeper:
		//新增操作
		return _zk.(*zookeeper.Zookeeper), nil
	default:
		return nil, fmt.Errorf("the zk:%s connect not exist", zkName)
	}
}

// GetOtherComponent 获取其他组件实例
func GetOtherComponent(ctx *gin.Context, name string) (interface{}, error) {
	_other, exist := ctx.Get(name)
	if !exist {
		return nil, fmt.Errorf("the %T:%s connect not exist", _other, name)
	}
	switch _other.(type) {
	case interface{}:
		//新增操作
		return _other.(interface{}), nil
	default:
		return nil, fmt.Errorf("the %T:%s connect not exist", _other, name)
	}
}

func GetK8S(ctx *gin.Context, k8sName string) (*k8s.Client, error) {
	_k8s, exist := ctx.Get(k8sName)
	if !exist {
		return nil, fmt.Errorf("the k8s:%s connect not exist", k8sName)
	}
	switch _k8s.(type) {
	case *k8s.Client:
		//新增操作
		return _k8s.(*k8s.Client), nil
	default:
		return nil, fmt.Errorf("the k8s:%s connect not exist", k8sName)
	}
}

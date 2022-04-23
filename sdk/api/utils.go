package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.benlai.work/go/ymir/storage/db"
)

// GetOrm 获取orm连接
func GetOrm(c *gin.Context, dbName string) (*db.Yorm, error) {
	idb, exist := c.Get(dbName)
	if !exist {
		return nil, fmt.Errorf("the db:%s connect not exist", dbName)
	}
	switch idb.(type) {
	case *db.Yorm:
		//新增操作
		return idb.(*db.Yorm), nil
	default:
		return nil, fmt.Errorf("the db:%s connect not exist", dbName)
	}
}

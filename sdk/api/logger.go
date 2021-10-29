package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/sdk/api/request"
	"strings"
)

// GetRequestLogger 获取上下文提供的日志
func GetRequestLogger(c *gin.Context) *logger.Helper {
	var log *logger.Helper
	l, ok := c.Get(request.LoggerKey)
	if ok {
		ok = false
		log, ok = l.(*logger.Helper)
		if ok {
			return log
		}
	}
	//如果没有在上下文中放入logger
	requestId := GenerateMsgIDFromContext(c)
	log = logger.NewHelper(logger.DefaultLogger).WithFields(map[string]interface{}{
		strings.ToLower(request.TrafficKey): requestId,
	})
	return log
}

// SetRequestLogger 设置logger中间件
func SetRequestLogger(c *gin.Context) {
	requestId := GenerateMsgIDFromContext(c)
	log := logger.NewHelper(logger.DefaultLogger).WithFields(map[string]interface{}{
		strings.ToLower(request.TrafficKey): requestId,
	})
	c.Set(request.LoggerKey, log)
}

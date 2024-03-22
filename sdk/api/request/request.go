package request

import (
	"github.com/geniussheep/ymir/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId, exist := c.Get(logger.TrafficKey)
	if requestId == "" || !exist {
		requestId = uuid.New().String()
		c.Set(logger.TrafficKey, requestId)
	}
	return requestId.(string)
}

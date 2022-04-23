package request

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	TrafficKey = "X-Request-Id"
	LoggerKey  = "_ymir-logger-request"
)

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(TrafficKey)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(TrafficKey, requestId)
	}
	return requestId
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.benlai.work/go/ymir/sdk/api/request"
)

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(request.TrafficKey)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(request.TrafficKey, requestId)
	}
	return requestId
}

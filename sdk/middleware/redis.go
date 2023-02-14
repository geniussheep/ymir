package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.benlai.work/go/ymir/sdk"
	"gitlab.benlai.work/go/ymir/sdk/config"
)

func WithContextRedis(c *gin.Context) {
	for redisName := range config.RedisConfig {
		c.Set(redisName, sdk.Runtime.GetRedis(redisName).WithContext(c))
	}
	c.Next()
}

package middleware

import (
	"github.com/geniussheep/ymir/sdk"
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/gin-gonic/gin"
)

func WithContextRedis(c *gin.Context) {
	for redisName := range config.RedisConfig {
		c.Set(redisName, sdk.Runtime.GetRedis(redisName).WithContext(c))
	}
	c.Next()
}

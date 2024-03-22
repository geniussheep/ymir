package middleware

import (
	"github.com/geniussheep/ymir/sdk"
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/gin-gonic/gin"
)

func WithContextZookeeper(c *gin.Context) {
	for zkName := range config.ZookeeperConfig {
		c.Set(zkName, sdk.Runtime.GetZookeeper(zkName))
	}
	c.Next()
}

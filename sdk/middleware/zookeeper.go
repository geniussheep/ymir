package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.benlai.work/go/ymir/sdk"
	"gitlab.benlai.work/go/ymir/sdk/config"
)

func WithContextZookeeper(c *gin.Context) {
	for zkName := range config.ZookeeperConfig {
		c.Set(zkName, sdk.Runtime.GetZookeeper(zkName))
	}
	c.Next()
}

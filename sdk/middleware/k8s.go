package middleware

import (
	"github.com/geniussheep/ymir/sdk"
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/gin-gonic/gin"
)

func WithContextK8S(c *gin.Context) {
	for k8sName := range config.K8SConfig {
		c.Set(k8sName, sdk.Runtime.GetK8S(k8sName))
	}
	c.Next()
}

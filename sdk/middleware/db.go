package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.benlai.work/go/ymir/sdk"
	"gitlab.benlai.work/go/ymir/sdk/config"
)

func WithContextDb(c *gin.Context) {
	for dbName := range config.DatabaseConfig {
		c.Set(dbName, sdk.Runtime.GetDb(dbName).WithContext(c))
	}
	c.Next()
}

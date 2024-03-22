package middleware

import (
	"github.com/geniussheep/ymir/sdk"
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/gin-gonic/gin"
)

func WithContextDb(c *gin.Context) {
	for dbName := range config.DatabaseConfig {
		c.Set(dbName, sdk.Runtime.GetDb(dbName).WithContext(c))
	}
	c.Next()
}

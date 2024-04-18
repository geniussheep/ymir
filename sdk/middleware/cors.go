package middleware

import (
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/gin-gonic/gin"
)

// CORS 跨域头设置
func CORS(c *gin.Context) {

	origin := c.Request.Header.Get("origin")
	if len(origin) <= 0 {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.HttpConfig.Cors.AllowOrigin)
	} else {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	}

	c.Writer.Header().Set("Access-Control-Allow-Credentials", config.HttpConfig.Cors.AllowCredentials)
	c.Writer.Header().Set("Access-Control-Allow-Headers", config.HttpConfig.Cors.AllowHeaders)
	c.Writer.Header().Set("Access-Control-Allow-Methods", config.HttpConfig.Cors.AllowMethods)
	c.Writer.Header().Set("Access-Control-Max-Age", config.HttpConfig.Cors.MaxAge)
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

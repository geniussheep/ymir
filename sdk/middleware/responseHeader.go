package middleware

import (
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/gin-gonic/gin"
)

// CustomResponseHeaders 设置自定义Response的头
func CustomResponseHeaders(c *gin.Context) {
	if config.HttpConfig.ResponseHeaders != nil {
		for k, v := range config.HttpConfig.ResponseHeaders {
			c.Writer.Header().Set(k, v)
		}
	}

	c.Next()
}

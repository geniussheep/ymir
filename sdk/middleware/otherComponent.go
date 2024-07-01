package middleware

import (
	"github.com/geniussheep/ymir/sdk"
	"github.com/gin-gonic/gin"
	"strings"
)

func WithContextOtherComponent(c *gin.Context) {
	for k, v := range sdk.Runtime.GetAllOtherComponent() {
		c.Set(strings.ToLower(k), v)
	}
	c.Next()
}

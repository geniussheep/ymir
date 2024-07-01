package middleware

import (
	"fmt"
	"github.com/geniussheep/ymir/sdk/pkg"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {

			if c.IsAborted() {
				c.Status(200)
			}
			switch errStr := err.(type) {
			case string:
				statusCode := http.StatusBadRequest
				p := strings.Split(errStr, "#")
				if len(p) == 3 && p[0] == "CustomError" {
					if _statusCode, _err := strconv.Atoi(p[1]); _err != nil {
						statusCode = _statusCode
					}
				}
				fmt.Println(
					time.Now().Format("2006-01-02 15:04:05"),
					"[ERROR]",
					c.Request.Method,
					c.Request.URL,
					500,
					c.Request.RequestURI,
					pkg.GetClientIP(c),
					err,
				)
				c.Status(statusCode)
				c.JSON(http.StatusBadRequest, gin.H{
					"code": statusCode,
					"msg":  errStr,
				})
			default:
				panic(err)
			}
		}
	}()
	c.Next()
}

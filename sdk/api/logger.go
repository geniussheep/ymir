package api

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/sdk/api/request"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// GetRequestLogger 获取上下文提供的日志
func GetRequestLogger(c *gin.Context) *logger.Helper {
	var log *logger.Helper
	l, ok := c.Get(request.LoggerKey)
	if ok {
		ok = false
		log, ok = l.(*logger.Helper)
		if ok {
			return log
		}
	}
	//如果没有在上下文中放入logger
	requestId := request.GenerateMsgIDFromContext(c)
	log = logger.NewHelper(logger.DefaultLogger).WithFields(map[string]interface{}{
		strings.ToLower(request.TrafficKey): requestId,
	})
	return log
}

// SetRequestLogger 设置logger中间件
func SetRequestLogger(c *gin.Context) {
	requestId := request.GenerateMsgIDFromContext(c)
	log := logger.NewHelper(logger.DefaultLogger).WithFields(map[string]interface{}{
		strings.ToLower(request.TrafficKey): requestId,
	})
	c.Set(request.LoggerKey, log)
}

func logRequest(log *logger.Helper, path string, raw string, req *http.Request) {
	var buf bytes.Buffer
	tee := io.TeeReader(req.Body, &buf)
	body, _ := ioutil.ReadAll(tee)
	req.Body = ioutil.NopCloser(&buf)

	if raw != "" {
		path = path + "?" + raw
	}

	log.Debugw("",
		"path", path,
		"header", req.Header,
		"body", string(body),
	)

}

func LoggerMiddleware(log *logger.Helper, debug bool, notlogged ...string) gin.HandlerFunc {

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {

		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		if debug {
			if _, ok := skip[path]; !ok {
				logRequest(log, path, raw, c.Request)
			}
		}

		c.Next()

		if _, ok := skip[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)
			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
			bodySize := c.Writer.Size()
			lv := logger.InfoLevel
			if len(errorMessage) > 0 {
				lv = logger.ErrorLevel
			}

			if raw != "" {
				path = path + "?" + raw
			}

			log.Logw(
				lv,
				"",
				"latency", latency,
				"path", path,
				"clientIP", clientIP,
				"method", method,
				"statusCode", statusCode,
				"errorMessage", errorMessage,
				"bodySize", bodySize,
			)
		}
	}
}

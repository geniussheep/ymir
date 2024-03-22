package middleware

import (
	"bytes"
	"github.com/geniussheep/ymir/logger"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func logRequest(log *logger.Helper, path string, raw string, req *http.Request) {
	var buf bytes.Buffer
	tee := io.TeeReader(req.Body, &buf)
	body, _ := ioutil.ReadAll(tee)
	req.Body = ioutil.NopCloser(&buf)

	if raw != "" {
		path = path + "?" + raw
	}

	log.Debug(
		"path", path,
		"header", req.Header,
		"body", string(body))
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

			logger.Log(
				lv,
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

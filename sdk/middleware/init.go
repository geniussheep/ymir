package middleware

import "github.com/gin-gonic/gin"

var (
	middlewares = make([]gin.HandlerFunc, 0)
)

func AppendDefault() {
	middlewares = append(middlewares, WithContextDb)
	middlewares = append(middlewares, WithContextRedis)
	middlewares = append(middlewares, CustomError)
	middlewares = append(middlewares, CORS)
}

func Append(middleware gin.HandlerFunc) {
	middlewares = append(middlewares, middleware)
}

func InitMiddleware(r *gin.Engine) {
	for _, mwf := range middlewares {
		r.Use(mwf)
	}
}

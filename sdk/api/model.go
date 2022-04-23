package api

import "github.com/gin-gonic/gin"

type RouterEntry struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

type RouterGroup struct {
	RouterPrefix string
	Routers      []RouterEntry
}

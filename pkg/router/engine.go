package router

import (
	"dreamland/pkg/ctrl"
	"github.com/gin-gonic/gin"
)

func Default() *gin.Engine {
	engine := gin.Default()
	engine.GET("/ping", ctrl.Example.Ping)

	return engine
}

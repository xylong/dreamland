package router

import (
	"dreamland/pkg/ctrl"
	"github.com/gin-gonic/gin"
)

func Default() *gin.Engine {
	engine := gin.New()
	engine.Use(recovery)
	engine.Use(responseHandler)

	engine.GET("/ping", ctrl.Example.Ping)
	engine.GET("/404", ctrl.Example.NotFound)
	engine.GET("/ok", ctrl.Example.OK)

	return engine
}

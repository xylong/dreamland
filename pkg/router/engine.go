package router

import (
	"dreamland/pkg/ctrl"
	"dreamland/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Default() *gin.Engine {
	engine := gin.New()
	engine.Use(middleware.Recovery)
	engine.Use(middleware.ResponseHandler)

	engine.GET("/ping", ctrl.Example.Ping)
	engine.GET("/404", ctrl.Example.NotFound)
	engine.GET("/ok", ctrl.Example.OK)

	engine.POST("/register", ctrl.User.Login)

	return engine
}

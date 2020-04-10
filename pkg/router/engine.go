package router

import (
	"dreamland/api/v1"
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

	engine.POST("/register", ctrl.User.Register)
	engine.POST("/login", ctrl.Authorization.Login)

	api := engine.Group("/api/v1")
	api.Use(middleware.JWT())
	{
		api.GET("/user", v1.User.Me)
	}

	return engine
}

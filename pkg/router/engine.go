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

	engine.GET("/ping", ctrl.Example.Ping)
	engine.GET("/404", ctrl.Example.NotFound)
	engine.GET("/ok", ctrl.Example.OK)

	engine.POST("/login", v1.Authorization.Login)

	api := engine.Group("/api")
	{
		V1 := api.Group("/v1")
		V1.POST("/register", v1.User.Store)
		V1.PUT("/authorizations/current", v1.Authorization.Current)

		V1.Use(middleware.JWT(), middleware.ResponseHandler)
		{
			V1.POST("/login", v1.Authorization.Login)
			V1.GET("/user", v1.User.Me)
		}
	}

	return engine
}

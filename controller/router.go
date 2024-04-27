package controller

import (
	"leong-01-goapp/controller/middleware"

	"github.com/gin-gonic/gin"
)

func route(engine *gin.Engine) {
	engine.Use(middleware.CorsMiddleware)
	rg := engine.Group("/api/v1")
	{
		rg.GET("/ping", HandlePing)
	}

}

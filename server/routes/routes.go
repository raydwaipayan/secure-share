package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raydwaipayan/secure-share/handlers"
	"github.com/raydwaipayan/secure-share/server/handlers"
)

func registerRoutes(router *gin.Engine) {
	g := router.Group("/file")

	g.POST("/submit", handlers.Submit)
	g.POST("/retrieve", handlers.Retrieve)
}

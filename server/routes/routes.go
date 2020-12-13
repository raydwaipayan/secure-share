package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/raydwaipayan/secure-share/server/handlers"
)

// RegisterRoutes takes a gin instance and registers the routes
func RegisterRoutes(router *gin.Engine) {
	g := router.Group("/file")

	g.POST("/put", handlers.Submit)
	g.GET("/get", handlers.Retrieve)
}

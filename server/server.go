package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raydwaipayan/secure-share/config"
	"github.com/raydwaipayan/secure-share/server/routes"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	conf := config.Load()

	// ping address
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Register routes
	routes.RegisterRoutes(r)

	// Listen and serve on :Port
	r.Run(fmt.Sprintf(":%v", conf.Port))
}

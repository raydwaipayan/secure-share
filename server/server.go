package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raydwaipayan/secure-share/config"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	conf := config.Load()

	// Listen and serve on :Port
	r.Run(fmt.Sprintf(":%v", conf.Port))
}

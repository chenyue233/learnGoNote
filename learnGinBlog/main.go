package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	initializeRoutes()
	r.Run(":8080")
}




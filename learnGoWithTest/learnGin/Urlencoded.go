package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LoginFrom struct {
	User     string `from:"user" binding:"required"`
	Password string `from:"password" binding:"required"`
}

func main(){
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var form LoginFrom
		if c.ShouldBindWith(&form,binding.XML) == nil{
			c.JSON(200,gin.H{"status":"you are logged in"})
		}else {
			c.JSON(401,gin.H{"status":"unauthorized"})
		}
	})
	r.Run(":8080")
}


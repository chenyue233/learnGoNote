package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main()  {
	r := gin.Default()
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		Page := c.DefaultQuery("page","0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id :%s;page:%s;name:%s;message:%s",id,Page,name,message)
	})
	r.Run(":8080")
}

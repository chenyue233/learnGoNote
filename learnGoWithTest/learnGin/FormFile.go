package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"path/filepath"
)

func main(){
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Static("/","./public")
	r.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		file,err := c.FormFile("file")

		if err != nil{
			c.String(http.StatusBadRequest,fmt.Sprintf("get from err:%s",err.Error()))
			return
		}
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file,filename);err != nil{
			c.String(http.StatusBadRequest,fmt.Sprintf("upload file err:%s",err.Error()))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
	})
	r.Run(":8080")

}

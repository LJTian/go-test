package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("user: %s  password: %s type : %s", username, password, types))
	})
	r.POST("/upload", func(c *gin.Context) {
		_, file, err := c.Request.FormFile("file")
		if err != nil {
			c.String(http.StatusInternalServerError, "上传图片失败！")
		}

		if file.Size > 1024*1024*2 {
			c.String(http.StatusOK, "文件太大了")
			return
		}
		dst := "png/" + file.Filename
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, file.Filename)
	})
	r.Run()
}

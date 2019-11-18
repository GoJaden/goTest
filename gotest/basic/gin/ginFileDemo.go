package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.Writer.Write([]byte("test success"))
		return
	})

	//上传文件
	r.POST("/upload", func(c *gin.Context) {
		r.MaxMultipartMemory = 100 << 20
		fd, err := c.FormFile("file")

		if err != nil {
			c.JSON(200, "上传文件失败"+err.Error())
			return
		}
		fmt.Printf("文件名:%s,文件大小:%d\n", fd.Filename, fd.Size)

		f, err1 := fd.Open()
		defer f.Close()
		if err1 != nil {
			c.JSON(1000, "上传文件失败"+err.Error())
			return
		}
		file, _ := os.Create("d:/" + fd.Filename)
		defer file.Close()
		_, err3 := io.Copy(file, f)
		fmt.Println(err3)
	})
	r.Run(":8081")
}

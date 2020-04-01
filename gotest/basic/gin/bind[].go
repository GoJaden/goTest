package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Per struct {
	Ids []int `form:"ids"`
}

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {

		fmt.Println("---")

		go func() {

			time.Sleep(time.Second*5)
			fmt.Println("++++")
		}()
		c.JSON(200, "hello")
	})
	r.Run()
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Per struct {
	Ids []int `form:"ids"`
}

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		p := new(Per)
		V := c.BindQuery(p)
		fmt.Println(V, p)
		c.JSON(200, nil)
	})
	r.Run()
}

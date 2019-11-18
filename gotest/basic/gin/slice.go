package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type OrderListByStatusRequest struct {
	UserID        uint64 `form:"userId"`
	ExamineStatus []int  `form:"examineStatus"`
	PageSize      int    `form:"pageSize"`
	PageNum       int    `form:"pageNum"`
}

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		args := new(OrderListByStatusRequest)
		if err := c.BindQuery(args); err != nil {
			fmt.Println(err)
		}
		c.JSON(200, args)
	})

	r.Run()

}

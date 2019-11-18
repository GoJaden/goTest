package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math"
)

//体检排期 请求
type ScheduleRequest struct {
	//机构id  必传
	AgencyID string `form:"agencyId" json:"agencyId"`
	SkuId    int    `form:"skuId" json:"skuId"`
}

func main() {

	fmt.Print("number:", math.IsNaN(9221120237041090561))
	g := gin.Default()
	g.GET("/bindQuery/:id2/*id", func(c *gin.Context) {
		id := c.Param("id")
		id2 := c.Param("id2")
		fmt.Println(id)
		fmt.Println(id2)
	})

	g.GET("/bindYesOrNo", func(context *gin.Context) {
		//如果不传，就是默认值
		req := new(ScheduleRequest)

		if err := context.BindQuery(req); err != nil {
			context.JSON(200, err)
		}
		context.JSON(200, req)

	})
	//不使用参数绑定，解析json数据
	g.POST("/json", func(c *gin.Context) {
		body := c.Request.Body
		fmt.Println(body)

		defer body.Close()
		b, _ := ioutil.ReadAll(body)
		user := new(Users)
		err := json.Unmarshal(b, &user)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, user)
	})

	g.POST("/bindJson", func(c *gin.Context) {
		user := new(Users)
		c.BindJSON(&user)
		c.JSON(200, user)

	})

	g.POST("/bindQuery", func(c *gin.Context) {
		test := new(Test)
		c.BindQuery(&test)
		c.JSON(200, gin.H{"user": test,
			"code": 1000})
	})

	g.GET("/queryMap", func(c *gin.Context) {
		m, err := c.GetQueryMap("map")
		if err == false {
			fmt.Println(err)
			return
		}
		fmt.Printf("%T,%s", m, m)
		c.JSON(200, m)
	})

	g.POST("/postQueryMap", func(c *gin.Context) {
		m, err := c.GetQueryMap("map")
		if err == false {
			fmt.Println(err)
			return
		}
		fmt.Printf("%T,%s", m, m)
		c.JSON(200, m)
	})

	g.POST("/postFormMap", func(c *gin.Context) {
		m := c.PostFormMap("map")
		fmt.Println(m)
		c.JSON(200, m)
	})

	g.POST("/postFormList", func(c *gin.Context) {
		list := c.PostFormArray("list")
		fmt.Println(list)
		c.JSON(200, list)
	})

	g.Run(":8082")
}

type Test struct {
	S1 string `form:"s1"`
	S2 string `form:"s2"`
}
type Users struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Age      int    `json:"age"`
	IsAllow  bool   `json:"isAllow"`
}

type ReqUserEntity struct {
	/*	UserId int ``*/
}

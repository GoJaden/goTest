package main

import (
	httputil "common-lib/http-util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Data struct {
	SkuIds []string `form:"skuIds",json:"skuIds"`
}

func main() {
	r := gin.Default()

	r.GET("/demo", func(context *gin.Context) {

		fmt.Println("12323")
	})

	r.POST("/test", func(context *gin.Context) {
		ids := new(Data)
		context.BindJSON(ids)
		a, _ := json.Marshal(ids)
		fmt.Println("-------", string(a))
		context.Writer.Write(a)
	})
	r.GET("/test2", func(context *gin.Context) {

		s := httputil.NewHttpClient()
		ss := new(Data)
		ss.SkuIds = append(ss.SkuIds, "1", "2")
		err := s.DoPostJson("http://localhost:8080/test", "", ss).ResponseParseStruct(ss)
		fmt.Println(err, ss)
	})
	type dd struct {
		D interface{} `form:"p"`
	}
	r.GET("/test3", func(context *gin.Context) {
		d := new(dd)

		context.BindQuery(d)

		fmt.Println(d)

		context.JSON(200, d)

	})

	r.Run()
}

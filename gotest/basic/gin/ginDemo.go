package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		V := c.Request.URL.Query()
		fmt.Println(V)
		c.JSON(200, nil)
	})
	/*api :=r.Group("/api")

		api.GET("/index/:id",index)
		api.POST("/user",addUser)
		api.PUT("/put",put)
		api.GET("/userinfo",bindUserinfo)
	r.GET("/index/:id",index)
	r.POST("/user",addUser)
	r.PUT("/put",put)
	r.DELETE("/delete",buildDelete)
	r.PATCH("/buildpatch", BuildPatch)
	r.HEAD("/buildhead", BuildHead)
	r.OPTIONS("/buildoptions", BuildOptions)*/
	r.Run()
}

/*func bindUserinfo(c *gin.Context) {
	user := new(ReqQueryUser)
	fmt.Println(user)

	if err :=c.Bind(&user);err!=nil{
		c.JSON(200,"绑定失败")
		return
	}
	fmt.Println("获取到用户ID：",user.UserId)
	c.JSON(200,user)
}

type ReqQueryUser struct {
	UserId int64 `form:"userId"`
}

func BuildOptions(c *gin.Context) {

	c.JSON(200,"Options请求的结果")
}*/

func BuildHead(c *gin.Context) {
	fmt.Println("开启head请求")
	c.JSON(200, "head请求的结果")

}

func BuildPatch(c *gin.Context) {
	c.JSON(200, "Patch请求的结果")
}

func buildDelete(c *gin.Context) {
	c.JSON(200, "buildDelete请求的结果")
}

func put(c *gin.Context) {
	/*userid :=c.Param("id")
	fmt.Println(userid)
	set := collecttion.NewHashSet(1,2,3)
	fmt.Println(set.Exist(userid))
	c.JSON(200,"lala")*/
}

//post 获取数据方式
func addUser(c *gin.Context) {
	id := c.PostForm("id")
	userid, _ := strconv.Atoi(id)
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := User{userid, username, password}
	fmt.Println(id, "-", username, "-", password)

	c.JSON(http.StatusOK, user)
}

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
type Student struct {
	//需要大写，不然反射获取不到数据
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func index(c *gin.Context) {
	//参数化路径
	id := c.Param("id")
	//stu := Student{20,"jiangjiu",23}
	//json格式化数据
	//data,_:=json.Marshal(stu)
	//获取get 请求参数
	fmt.Println(c.Query("id"))
	fmt.Println(c.Query("code"))
	c.String(http.StatusOK, id)
}

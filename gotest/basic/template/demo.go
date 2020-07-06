package main

import (
	"html/template"
	"net/http"
)

//使用模板渲染方式

func main() {
	/*g := gin.Default()
	g.Handle("GET","/test", func(context *gin.Context) {
		m := map[string]string{"jaden":"laoda","xxl":"xiaoer"}
		context.HTML(200,"data",&m)
	})
	g.Run(":8080")*/
	a := new(A)
	http.ListenAndServe(":8080", a)
}

type A int

func (a A) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("gotest/basic/template/templates/index.html")
	if err != nil {
		resp.Write([]byte("模板渲染失败" + err.Error()))
		return
	}
	tmp.Execute(resp, "jaden is good")

}

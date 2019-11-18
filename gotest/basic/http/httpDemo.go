package main

import (
	"fmt"
	"net/http"
)

func main() {
	//缺点：默认的需要自己去判断请求方式 的
	http.HandleFunc("/index", test)
	http.HandleFunc("/login", login)
	//第二个参数 是路由器的功能，为nil，默认是DefaultServerMux
	http.ListenAndServe(":8080", nil)

}

func login(w http.ResponseWriter, r *http.Request) {

	//cookie用法
	/*r.Cookie("jaden")*/
	/*cookie := http.Cookie{
		Name:"jaden",
		Value:"testCookie",
		Path:"/",
	}
	http.SetCookie(w,&cookie)*/
	r.ParseForm()
	if r.Method == "GET" {
		w.Write([]byte("不支持get请求登录"))
		return
	}
	for po, val := range r.PostForm {
		fmt.Printf("位置%s,值%s\n", po, val)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求方式是", r.Method)
	//解析url传递的参数，对于post则解析响应包的主体
	r.ParseForm()
	//需要在之前调用r.ParseForm（）才可以，详情见注解
	fmt.Println(r.Form)
	var buffer []byte = make([]byte, 1024)
	r.Body.Read(buffer)
	fmt.Println(string(buffer))
	fmt.Println(r)
	fmt.Println(w)
}

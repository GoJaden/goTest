package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Print(req)
	v := req.PostFormValue("TEST")
	v1 := req.PostFormValue("TEST")
	v2 := req.PostFormValue("TEST")
	v3 := req.PostFormValue("TEST")
	fmt.Print(v)
	fmt.Print(v1)
	fmt.Print(v2)
	fmt.Print(v3)
}

func main() {
	//url 与  处理器 匹配
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务端开启失败", err)
	}
}

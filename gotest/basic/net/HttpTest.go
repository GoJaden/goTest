package main

import (
	"fmt"
	"go.uber.org/atomic"
	"net/http"
	"sync"
	"time"
)

var count atomic.Int64
var mutx sync.Mutex

func indexHandler(w http.ResponseWriter, req *http.Request) {
	mutx.Lock()

	count.Inc()
	time.Sleep(time.Second / 2)
	mutx.Unlock()
	fmt.Printf("处理了%d个,%f\n", count.Load(), 5e7)
	http.Redirect(w, req, "http://www.baidu.com", 301)

	return
}

func main() {
	//url 与  处理器 匹配
	http.HandleFunc("/a", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务端开启失败", err)
	}
}

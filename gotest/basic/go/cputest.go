package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	Test()
}

//测试在单核情况下
func Test() {
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second / 5)
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second * 10)
}

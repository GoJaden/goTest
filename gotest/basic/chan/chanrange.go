package main

import (
	"fmt"
	"time"
)

func main() {

	test()

}

func send() chan int {
	tt := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("开始发送数据")
		tt <- 2
		close(tt)
	}()
	return tt
}

func test() {

	tt := send()

	for dat := range tt {
		fmt.Println(dat)
	}
	fmt.Println("----")

}

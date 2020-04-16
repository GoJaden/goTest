package main

import (
	"fmt"
	"time"
)

func timeMore(ch chan string) {

	// 执行前先注册，写不进去就会阻塞
	ch <- "任务"

	fmt.Println("模拟耗时操作")
	time.Sleep(time.Second) // 模拟耗时操作

	// 任务执行完毕，则管道中销毁一个任务
	data := <-ch
	fmt.Println(data)
}

//可以用于限流
func main() {
	print("fdfd\n")
	ch := make(chan string, 5)

	// 开启100个协程
	for i := 0; i < 100; i++ {
		go timeMore(ch)
	}

	for {
		time.Sleep(time.Second * 5)
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {

	/*ch := make(chan int,2)

	// 首先，我们实现并执行一个匿名的超时等待函数
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second*3) // 等待1秒钟
		timeout <- true
	}()
	// 然后我们把timeout这个channel利用起来
	select {
	case <-ch:
	// 从ch中读取到数据
	case <-timeout:
		// 一直没有从ch中读取到数据，但从timeout中读取到了数据
		fmt.Println("超时等待")
	}*/

	timer

	/*
		ch := make(chan int,2)

			ch <- 0
			fmt.Println(len(ch))*/
	/*data1 := <-ch
	fmt.Println("data1:",data1)*/

	/*go func() {
		ch <- 0
		fmt.Println("发送数据数据")
	}()
	data := <-ch
	fmt.Println("data:",data)*/

	/*go test()

	var input string
	fmt.Scanln(&input)*/

}

var tik int

func test() {

	for {
		tik++
		fmt.Println("tik:", tik)
		time.Sleep(time.Second * 1)
	}
}

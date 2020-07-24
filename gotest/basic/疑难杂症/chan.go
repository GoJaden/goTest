package main

import (
	"fmt"
	"time"
)

//无缓冲的chan测试

func NoBufferChanTest() {
	ch1 := make(chan int)
	fmt.Println(ch1)
	go func() {
		fmt.Println("开始写入")
		//无缓冲的chanle，写入 跟 读取都会发生阻塞，会发生goroutine泄露问题
		ch1 <- 1

		fmt.Println("写入完成")
	}()

	go func() {
		fmt.Println(<-ch1)
		fmt.Println("再次读取", <-ch1)
		fmt.Println("读取完成")
	}()
	fmt.Println("main1")
	time.Sleep(time.Second * 3)
	fmt.Println("main2")
}

/*
	关闭没有make的chan会报panic
	var cChan chan int
	close(cChan)

*/

func main() {
	NoBufferChanTest()

}

//空通道的测试
func nullChanTest(nullChan chan int) {
	go func() {
		fmt.Println("空通道测试", "通道长度", len(nullChan))
		for {
			select {
			case res := <-nullChan:
				{
					fmt.Println("获取到:", res)
				}
			default:
				fmt.Println("没有从通道中获取到值")
			}
		}

	}()
}

//长度为1 的通道测试
func defaultChanTest(defaultChan chan int) {

}

//长度不为1的通道测试
func lenChanTest(lenChan chan int) {

}

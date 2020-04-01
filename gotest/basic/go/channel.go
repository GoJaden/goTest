package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	test2()
}

func test2() {
	var slice []int = []int{}
	data, err := json.Marshal(slice)
	fmt.Println(string(data), err)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c1 := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("准备发送数据")
		c1 <- 1
		wg.Done()
	}()

	da, ok := <-c1
	if ok {
		fmt.Println("接收到的数据：", da)
	}

	wg.Wait()
	//time.Sleep(time.Second*5)

}

func funcMui(x, y int) (sum int, a error) {
	return x + y, nil
}
func test1() {
	runtime.GOMAXPROCS(2)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)

	case value := <-string_chan:
		fmt.Println(value)

	}

	time.Sleep(time.Second * 4)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	Test()
}

// new(Type)*Type
// make(Type,size)Type

// new 用于任何类型，但是对于chan 、map、slice的话，并不能初始化引用中的值
//make 的话，会初始化引用类型的变量还有地址

func Test() {
	newSlice := new([]int)
	(*newSlice)[1] = 1
	fmt.Println(newSlice)
	newChan := new(chan int)
	go func() {
		fmt.Println(*newChan)
		*newChan <- 1
		fmt.Println(<-*newChan)
	}()

	time.Sleep(time.Second * 3)
}

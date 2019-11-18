package main

import (
	"fmt"
	"time"
)

/**
*首字母大写，表示public类型
*通过 包名.方法名 调用相对应的func
*这个地方注意配置项目的gopath路径
 */
func Hello() {
	str := "test1"
	data := []byte(str)

	fmt.Println(len(data))
}

var i int

func main() {

	for count := 0; count < 100; count++ {
		go test1()
	}
	time.Sleep(time.Second * 5)
	fmt.Println(i)
}

func test1() {
	go func() {
		for j := 0; j < 100; j++ {
			i++
		}

	}()
}

package main

import (
	"fmt"
	"testing"
)

//defer 跟 panic的执行顺序
//defer先执行，之后再执行panic
func main() {

	println("11111")
	fmt.Println("22222")
	m := new(MyMap)
	fmt.Println("%T", m)

	Test1(nil)
}
func Test1(t *testing.T) {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

type MyMap map[int]string

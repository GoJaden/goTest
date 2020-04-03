package main

import (
	"fmt"
	"github.com/CatBluePoor/moles"
)

func main() {

	defer moles.Release() // 释放协程池
	moles.Submit(Test)    // 提交任务
}
func Test() {
	fmt.Println("test")
}

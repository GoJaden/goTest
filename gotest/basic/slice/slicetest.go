package main

import (
	"fmt"
	"time"
)

func main() {
	//切片复制
	/*a := make([]int,0)
	a = append(a,19,15)
	fmt.Println(a[:])
	sClone := append(a[:0:0],a...)
	fmt.Println(sClone)
	b := make([]int,len(a))
	copy(b,a)
	fmt.Println(b)
	*/
	test2()
	time.Sleep(time.Second * 3)

}

func _() {
	fmt.Println("--")
}

//闭包特点 --所有goroutine都是访问外部的i

func test2() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

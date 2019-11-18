package main

import "fmt"

func test() {
	defer func() {
		fmt.Println("start....")
		if re := recover(); re != nil {
			fmt.Println("捕获到了异常,", re)
		}
		fmt.Println("end....")

	}()
	panic("发生了未知错误")

}

func main() {

	test()
}

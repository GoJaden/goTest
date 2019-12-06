package main

import (
	"fmt"
	"unicode/utf8"
)

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
	data := "♥"
	fmt.Println(utf8.RuneCountInString(data)) //prints: 3
	//test()
}

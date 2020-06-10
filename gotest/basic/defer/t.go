package main

import "fmt"

func main() {
	a := "你好bj!"
	fmt.Println(len("你好bj!"))
	for _, value := range a {
		fmt.Println(string(value))
	}

	/*
		先defer的后执行
		recover后输出panic中的信息
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
		} else {
			fmt.Print("no")
		}
	}()
	defer func() {
		panic("1111111111111")
	}()
	panic("22222222222")

}

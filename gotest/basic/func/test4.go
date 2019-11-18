package main

import "fmt"

func main() {
	fmt.Println(Add(1, 2))

}

//匿名函数
var Add = func(a, b int) int {
	return a + b
}

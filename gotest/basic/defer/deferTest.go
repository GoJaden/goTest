package main

import "fmt"

func main() {

	fmt.Println(test())

}

//defer带函数的执行，defer执行在返回赋值后，返回前执行
func test() int {
	res := 7
	defer func(res int) {
		fmt.Println("start,res=", res)
		res++
		fmt.Println("end,res=", res)
	}(res)
	return res
}

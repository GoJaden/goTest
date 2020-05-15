package main

import "fmt"

func main() {

	fmt.Println(f3())

	fmt.Println("res:", doSomething())

}
func doSomething() (rev int) {
	defer func() {
		rev++
	}()

	return 5
}

func f3() (r int) {
	defer func() {
		r++
	}()
	return 0
}

//defer执行顺序
// 返回值赋值
//defer
//return

func test() int {

	res := 7
	defer func() {
		fmt.Println("start,res=", res)
		res++
		fmt.Println("end,res=", res)
	}()

	res += 5
	return res
}

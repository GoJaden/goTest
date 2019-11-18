package main

import "fmt"

func main() {
	//函数类型的变量
	ad := add
	fmt.Printf("%d\n", ad(11, 2))

	fmt.Printf("%d\n", calc(12, 21, add))
	a := myFunc(add)
	b := myFunc(mod)

	fmt.Printf("a=%d,b=%d\n", a, b)
	a.Calc()
}

func calc(a int, b int, t func(a, b int) int) int {
	return t(a, b)
}

func (m myFunc) Calc() {
	fmt.Printf("开始计算...%T", m)
}

func add(a, b int) int {
	return a + b
}
func mod(a, b int) int {
	return a / b
}

type myFunc func(a, b int) int

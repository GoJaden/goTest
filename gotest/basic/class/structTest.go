package main

import "fmt"

type A struct {
	sx string
}

func (a *A) testA() {
	fmt.Println("a的方法")
}

//如果是匿名字段，则类似于 extend
type B struct {
	A
}

//如果是一个实际的字段，类似于组合
type C struct {
	a A
}

func main() {

	a := new(B)
	a.testA()

	b := new(C)
	c := b.a
	fmt.Printf("c:%v,ct:%T", c, c)
}

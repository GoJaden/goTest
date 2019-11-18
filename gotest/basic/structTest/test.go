package main

import "fmt"

//继承方式
type A struct {
	B
}

//组合方式
type C struct {
	BB B
}

type B struct {
	id int
}

func (b B) testB() {
	fmt.Println("hello")
}

func main() {
	a := new(A)
	a.testB()

	c := new(C)
	b := new(B)
	c.BB = *b
	c.BB.testB()

}

package main

import "fmt"

type Human interface {
	Say() string
}

type Man struct {
}

func (m *Man) Say() string {
	return "man"
}

func IsNil(h interface{}) bool {
	return h == nil
}

func main() {
	var a interface{}
	var b *Man
	var c *Man
	var d Human
	var e interface{}
	a = b
	e = a
	fmt.Println(a == nil)                  // (1) false
	fmt.Printf("%T,%v\n", e, e == nil)     // (2) false
	fmt.Printf("%T,%T,%v\n", a, c, a == c) // (3) true
	fmt.Println(a == d)                    // (4) false
	fmt.Println(c == d)                    // (5) false
	fmt.Println(e == b)                    // (6) true
	fmt.Println(IsNil(c), c)               // (7) true
	fmt.Println(IsNil(d))                  // (8) true
}

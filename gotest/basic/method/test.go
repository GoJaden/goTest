package main

import "fmt"

type User struct {
	Id int
}

func (u User) test1() {
	fmt.Println(u, "+++")
}
func (u *User) test2() {
	fmt.Println(u, "---")
}

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
	fmt.Println(c == nil)

	fmt.Println(a == nil)
	// (1) false
	// a是eface类型，_type指向的是*Man类型，
	// data指向的是nil，所以此题为false

	fmt.Println(e == nil)
	// (2) false
	// 同理，e为eface类型，_type也是指向的*Man类型

	fmt.Println(a == c)
	// (3) true
	// a的_type是*Man类型，data是nil
	// c的data也是nil

	fmt.Println(a == d)
	// (4) false
	// a为eface类型，d为iface类型，而且d的itab指向的是nil，data也是nil
	// 因为d没有具体到哪种数据类型

	fmt.Println(c == d)
	// (5) false
	// c和d其实是两种不同的数据类型

	fmt.Println(e == b)
	// (6) true
	// 分析见(4)

	fmt.Println(IsNil(c))
	// (7) false
	// c是*Man类型，以参数的形式传入IsNil方法
	// 虽然c指向的是nil，但是参数i的_type指向的是*Man，所以i不为nil

	fmt.Println(IsNil(d))
	// (8) true
	// d没有指定具体的类型，所以d的itab指向的是nil，data也是nil
}

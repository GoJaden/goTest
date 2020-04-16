package main

import "fmt"

const (
	a = iota
	b = iota
)
const (
	name = "menglu"
	c    = iota << 1
	d    = iota
	e    = iota
)

func main() {
	fmt.Println([...]string{"1", "2"})
	//fmt.Println([]string{"1"} == []string{"1"})
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

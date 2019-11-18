package main

import "fmt"

type Father struct {
	*int
}

func main() {
	f := Father{}
	var a int = 1
	var b *int = &a
	f.int = b
	fmt.Println(f)
	test(f)
	fmt.Println(f)
}

func test(father Father) {
	var c int = 1
	var d *int = &c
	father.int = d
	fmt.Println(father)
}

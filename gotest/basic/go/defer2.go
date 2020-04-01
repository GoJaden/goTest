package main

import "fmt"

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func main() {
	a, b := test(100)
	a()
	b()

	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}
	fmt.Println(intmap)

	var x *int = nil
	Foo(x)
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
		fmt.Println("-", t)
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func(i int) {
		t += i
		fmt.Println("--", t)
	}(i)
	return 3

}

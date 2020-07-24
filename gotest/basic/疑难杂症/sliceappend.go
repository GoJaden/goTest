package main

import "fmt"

//参数传递为值传递，但是切片底层指向的切片的地址是一样的
func add_value(a []int, b int) {
	fmt.Printf("a参数的地址是%p,a的长度是%d,a的容量是%d\n", &a, len(a), cap(a))
	a = append(a, b, b)
	//a 的长度为2，容量为2
	a[0] = 123
	fmt.Println("append后的值", a)
	fmt.Printf("a append之后的值%p，长度是%d，容量是%d\n", &a, len(a), len(a))
}

func main() {
	//da 的切片长度为1，容量为2
	da := make([]int, 1, 2)
	da[0] = 1
	fmt.Printf("开始时da的地址%p,da的长度是%d,容量是%d\n", &da, len(da), cap(da))
	add_value(da, 3)
	fmt.Printf("结束时da的地址%p,da的长度是%d,容量是%d\n", &da, len(da), cap(da))
	fmt.Println(da)
}

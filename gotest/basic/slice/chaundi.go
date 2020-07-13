package main

import "fmt"

func Demo(slice []int) {
	slice[0] = 2
	fmt.Println("函数中结果：", slice)
}

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

const (
	xx = iota
	yy
)

func main() {
	fmt.Println(x, y, z, k, p, xx, yy)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:6] //3456  len=4 cap=6
	fmt.Println(len(s1), cap(s1))
	s2 := s1[3:6]
	fmt.Println(len(s2), cap(s2)) //len=3 cap=3

	fmt.Println("s1==", s1) // 3 4 5 6
	fmt.Println("s2==", s2) // 678

	abc := make([]int, 2)
	abc = append(abc, 1, 2, 3)
	fmt.Println(abc)

	msg := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//cap = max - lowmin
	sli1 := msg[2:3:4]
	//cap = 7
	sli2 := msg[2:3:7]
	fmt.Println(sli1, len(sli1), cap(sli1))
	fmt.Println(sli2, len(sli2), cap(sli2))

	//定义一个切片
	slice := make([]int, 10, 10)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	Demo(slice)
	fmt.Println("定义中结果：", slice)
}

package main

import "fmt"

func main() {
	dst := make([]int, 10)

	copy(dst, []int{1, 2, 3})
	fmt.Print(len(dst))
	//从指定位置开始复制元素
	copy(dst[3:], []int{4, 5, 6})
	fmt.Println(dst)
}

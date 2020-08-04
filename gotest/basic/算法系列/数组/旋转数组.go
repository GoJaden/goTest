package main

import "fmt"

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 5)
}

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
*/
/*  需要使用空间复杂度为O（1）的方法，使用到了数组翻转的思想，同时需要控制翻转位置*/
func rotate(dta []int, key int) {
	//先整体翻转
	reverse(dta)

	//再翻转指定位置的元素
	reverse(dta[:key%len(dta)])
	reverse(dta[key%len(dta):])

	fmt.Println(dta)
}

func reverse(data []int) {
	if len(data) == 1 {
		return
	}
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-i-1] = data[len(data)-i-1], data[i]
	}
}

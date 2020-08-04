package main

import "fmt"

//给定两个数组，编写一个函数来计算它们的交集。
func main() {
	a := []int{1, 2, 5, 6, 7, 8, 12, 12, 14}
	b := []int{1, 2, 3, 4, 6, 8, 9, 23}
	fmt.Println(jiaoji(a, b))
}

//如果给定的数组已经排好序呢
//双指针解法，相等同时后移，不等小的指针后移
func jiaoji(a, b []int) []int {

	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	poiA := 0
	poiB := 0
	result := make([]int, 0)
	for poiA < len(a) && poiB < len(b) {
		if a[poiA] == b[poiB] {
			result = append(result, a[poiA])
			poiA++
			poiB++
		} else if a[poiA] > b[poiB] {
			poiB++
		} else {
			poiA++
		}
	}
	return result
}

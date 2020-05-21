package main

import "fmt"

func main() {
	t := []int{45, 43, 343, 12, 43, 656, 656, 76, 34, 45, 4554, 65, 345, 543, 675, 231}
	fmt.Println(t)
	fmt.Println("插入排序", MergeSort(t, 0, len(t)-1))
}

/**
思想:将每个数组分成两部分,对应位置分别比较,交换位置
    递归遍历到长度为1,所有位置都进行了比较

*/

func Swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func MergeSort(arr []int, left, right int) []int {
	if len(arr) <= 1 {
		return arr
	}

	start := left
	end := right
	mid := (start + end) / 2
	midIndex := mid
	for left < mid && right > mid {
		if arr[left] > arr[mid] {
			Swap(arr, left, midIndex)
		}
		left++
		right--
		midIndex++
	}
	fmt.Println(arr)
	arr = MergeSort(arr[left:mid], start, mid)
	arr = MergeSort(arr[mid:], mid, end)
	return arr
}

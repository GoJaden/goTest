package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 2, 4, 5, 6, 7, 10, 14, 16, 17, 122}
	fmt.Println(BinarySearch1(data, 122, len(data)))
	fmt.Println(binarySearch([]int{1, 2, 2, 4, 5, 6, 7, 10, 14, 16, 17, 122}, 17, 0, 11))

}

//for循环遍历的版本
func BinarySearch1(data []int, t, len int) int {
	var mid, low, hight int
	hight = len - 1
	for low <= hight {
		mid = (low + hight) / 2
		if data[mid] == t {
			return mid
		}
		if data[mid] > t {
			hight = mid - 1
		}
		if data[mid] < t {
			low = mid + 1
		}
	}
	return -1

	return -1
}

//递归遍历的版本
func binarySearch(data []int, t, low, high int) int {
	mid := (high + low) / 2
	if data[mid] == t {
		return mid
	}
	if data[mid] > t {
		return binarySearch(data, t, low, mid-1)
	}
	if data[mid] < t {
		return binarySearch(data, t, mid+1, high)
	}
	return -1
}

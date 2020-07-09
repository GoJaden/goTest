package main

import "fmt"

func main() {
	t := []int{45, 43, 343, 12, 43, 656, 656, 76, 34, 45, 4554, 65, 345, 543, 675, 231}
	fmt.Println(t)
	fmt.Println("归并排序", mergeSort(t))
}

/**
思想:将每个数组分成两部分,对应位置分别比较,交换位置
    递归遍历到长度为1,所有位置都进行了比较

*/
func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2
	left := mergeSort(data[:mid])
	right := mergeSort(data[mid:])
	result := merge(left, right)
	return result
}

func merge(left []int, right []int) (result []int) {
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

/*func Swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func MergeSort(arr []int, left, right int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := (left + right) / 2
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
	arr = MergeSort(arr[left:mid], left, mid)
	arr = MergeSort(arr[mid:], mid, right)
	return arr
}
*/

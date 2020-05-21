package main

import "fmt"

func main() {
	t := []int{45, 43, 343, 12, 43, 656, 656, 76, 34, 45, 4554, 65, 345, 543, 675, 231}
	fmt.Println(t)
	fmt.Println("插入排序", InsertSort(t))

}

/*
简单插入排序思想:
遍历数组的每一个值，跟前面已经排序的数组进行比较
一直到比前面的位置小的时候，结束移动，循环一轮结束

效率: 时间复杂度 最差O(n*n)  最好O(n*n)
	  空间复杂度 O(1)

优点： 比较直观,不需要分配内存,内存消耗低

缺点： 每次都是双重遍历，效率低
*/

func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for tmpIndex := 1; tmpIndex < len(arr); tmpIndex++ {
		tmp := arr[tmpIndex]
		index := tmpIndex
		for index >= 1 && tmp <= arr[index-1] {
			arr[index-1], arr[index] = arr[index], arr[index-1]
			index--
		}
	}
	return arr
}

package main

import "fmt"

func main() {

	fmt.Println(SelectSort([]int{12, 3, 443, 5, 6, 2, 11, 32, 65, 7787, 8}))
}

/*选择排序，选择一个基数，每次都当成最大或者最小值*/

func SelectSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	for i := 0; i <= len(data)-1; i++ {
		minData := data[i]
		for j := i + 1; j < len(data); j++ {
			if minData > data[j] {
				minData = data[j]
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}

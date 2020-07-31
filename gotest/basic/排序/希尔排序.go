package main

import "fmt"

func main() {

	data1 := [][]int{{1, 23, 4}, {3, 2, 3}}
	fmt.Println(len(data1))
	fmt.Println(data1[1])

	data := []int{323, 43, 54, 13, 54, 65, 76, 34645, 65, 87, 432, 43, 1}
	fmt.Println(ShellSort(data))
}

/*希尔排序通过选取gap点，每次比较两个元素的方式，达到效率的提升*/
func ShellSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	gap := len(data) / 2

	for rindex, rd := range data[gap:] {
		if data[rindex] > rd {
			data[rindex], data[rindex+gap] = data[rindex+gap], data[rindex]
		}
	}
	fmt.Println("++", data)
	ShellSort(data[:gap])
	ShellSort(data[gap:])
	return data
}

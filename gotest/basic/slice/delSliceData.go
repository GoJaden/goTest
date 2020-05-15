package main

import "fmt"

func main() {

	// 删除切片中的某个元素

	// 删除切片中指定元素
	{
		items := []int{1, 2, 4, 2, 3, 0}
		deleteItem := 2
		for i := 0; i < len(items); i++ {
			if items[i] == deleteItem {
				items = append(items[:i], items[i+1:]...)
				i--
			}
		}
		fmt.Println(items) // output: [1, 4, 3, 0]
	}

	// 删除找到的第一个元素
	{
		// 切片比较大的话，还是用普通的 for 循环比较好
		items := []int{1, 2, 4, 2, 3, 0}
		deleteItem := 2
		for i, item := range items {
			// 找到要删除的第一个元素，删除并退出循环
			if item == deleteItem {
				items = append(items[:i], items[i+1:]...)
				break
			}
		}
		fmt.Println(items) // output: [1, 4, 2, 3, 0]
	}

}

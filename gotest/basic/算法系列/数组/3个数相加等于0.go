package main

import (
	"fmt"
	"sort"
)

/*给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。注意：答案中不可以包含重复的三元组。
示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]*/

func main() {
	dat := []int{-1, 0, 1, 2, -1, -4}
	sort.Slice(dat, func(i, j int) bool {
		return dat[i] < dat[j]
	})
	fmt.Println(ThreeNums(dat))
}

func ThreeNums(data []int) [][3]int {
	if len(data) <= 3 {
		return nil
	}
	result := make([][3]int, 0)
	for poi, res := range data {
		// 需要获取的两个相加的数
		num := 0 - res
		//定义左右指针范围
		l := poi + 1
		r := len(data) - 1
		if l > 0 && r < len(data)-1 {
			for l != r {
				if data[l]+data[r] == num {
					fmt.Println("i:", l, ",r:", r)
					result = append(result, [][3]int{{poi, r, l}}...)
					for l < r && data[l] == data[l+1] {
						l++
					}
					for l < r && data[r] == data[r-1] {
						r--
					}
					l++
					r--
				} else if data[l]+data[r] > num {
					l++
				} else {
					r++
				}
			}

		}

	}

	return result
}

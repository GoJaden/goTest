package main

import "fmt"

func main() {
	t := []int{45, 43, 343, 12, 43, 656, 656, 76, 34, 45, 4554, 65, 345, 543, 675, 231}
	fmt.Println(t)
	fmt.Println(QuickSort(t))
}

/*
简单快速排序思想:
1.取一个元素,将比它大的放在一个big数组中,比它小的放在另一个small数组中
2.迭代big跟small数组,对其中的每一个数组重复操作，一直到len(arr)<=1
3.返回排序数组的结果

效率: 时间复杂度 最坏O(n*n)  最好O(n*log(n))
      空间复杂度 最坏O（n）  平均O(log(n))

优点： 直观,速度快

缺点: 递归开辟内存,数据大，内存爆炸
*/

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	//待比较元素
	tmp := arr[0]
	big, small, equal := []int{}, []int{}, []int{}
	equal = append(equal, tmp)
	for k, v := range arr {
		if k == 0 {
			continue
		}
		if v > tmp {
			big = append(big, v)
		} else if v < tmp {
			small = append(small, v)
		} else {
			equal = append(equal, v)
		}
	}
	big = QuickSort(big)
	small = QuickSort(small)

	arr = append(small, equal...)
	arr = append(arr, big...)
	return arr
}

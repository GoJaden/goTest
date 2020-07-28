package main

import "fmt"

func main() {
	t := []int{45, 43, 343, 12, 43, 656, 656, 76, 34, 45, 4554, 65, 345, 543, 675, 231}
	fmt.Println(t)
	fmt.Println(QuickSort(t))

	t1 := []int{45, 43, 343, 12, 43, 656, 656, 76, 34, 45, 4554, 65, 345, 543, 675, 231}
	fmt.Println(DiGuiQuickSort(t1))

	fmt.Println("---")
	DoublePointerQuickSort(t, 0, len(t)-1)
	fmt.Println(t)
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

//todo  有待进一步考究
func DoublePointerQuickSort(data []int, left, right int) {
	if left < right {
		parResult := partition(data, left, right)
		DoublePointerQuickSort(data, left, parResult-1)
		DoublePointerQuickSort(data, parResult+1, right)
	}
}

var i = 0

//左右指针的方式方式实现快速排序  https://blog.csdn.net/qq_36528114/article/details/78667034?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.channel_param&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.channel_param
func partition(data []int, left, right int) int {
	begin := left
	end := right
	key := left
	for begin < end {
		for begin < end && data[begin] <= data[key] {
			begin += 1
		}

		for begin < end && data[end] >= data[key] {
			end -= 1
		}
		data[begin], data[end] = data[end], data[begin]
	}
	i++
	fmt.Printf("第%v次计算，结果是%v\n", i, data)
	return begin

}

//递归方式快排
func DiGuiQuickSort(data []int) []int {
	mid := len(data) / 2
	if len(data) <= 1 {
		return data
	}
	var leftData, midData, rightData []int
	for i := 0; i < len(data); i++ {
		if data[i] > data[mid] {
			rightData = append(rightData, data[i])
		} else if data[i] == data[mid] {
			midData = append(midData, data[i])
		} else {
			leftData = append(leftData, data[i])
		}
	}
	rightData = DiGuiQuickSort(rightData)
	leftData = DiGuiQuickSort(leftData)

	result := append(rightData, midData...)
	result = append(result, leftData...)
	return result
}

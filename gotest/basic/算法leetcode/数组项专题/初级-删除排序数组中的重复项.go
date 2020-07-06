package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 3, 3, 4, 4}))
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var first, last int
	for i := 1; i < len(nums); i++ {
		first = i
		last = i - 1
		for nums[first] == nums[last] && last < len(nums)-1 {
			nums[last] = nums[last+1]
			last++
		}
		first++
		last++
	}
	fmt.Println(nums)
	return len(nums)
}

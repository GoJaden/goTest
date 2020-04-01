package main

import "fmt"

func main() {

	fmt.Println(judgeOrderStatus([]int{2,2,2}))
}

//判断体检订单状态(区分 已预约跟已到检)
func judgeOrderStatus(status []int) int {
	if status == nil{
		return 0
	}

	statusMap := make(map[int]int,0)
	for _,s := range status{
		_ ,ok := statusMap[s]
		if !ok{
			statusMap[s] = s
		}
	}
	if len(statusMap) == 1{
		for key,_ := range statusMap{
			return key
		}
	}else if len(statusMap) >1{
		if _,ok := statusMap[1];ok{
			return 1
		}
	}
	return 2
}

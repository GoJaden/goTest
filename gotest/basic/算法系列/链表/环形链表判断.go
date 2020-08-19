package main

import (
	"fmt"
	"goProject/gotest/basic/算法系列/链表/domain"
	"time"
)

func main() {

	var mm chan int = make(chan int)
	close(mm)
	select {
	case data := <-mm:
		{
			fmt.Println("-", data)
		}
	}

	//用于监听完成事件
	doneChan := make(chan struct{}, 1)
	go func() {
		time.Sleep(time.Second * 3)
		doneChan <- struct{}{}
		fmt.Println("已经执行了")
	}()

	for {
		select {
		case doneChan <- struct{}{}:
			{
				fmt.Println("塞数据了")
			}
		case doneChan <- struct{}{}:
			{
				fmt.Println("塞数据了111")
			}
		case _, ok := <-doneChan:
			{
				if ok {
					fmt.Println("ok~~~~")
				}
				return
			}
		case <-time.After(time.Second * 2):
			{
				fmt.Println("超时")
				break
			}
		default:
			fmt.Println("default")

		}
		break
	}
	time.Sleep(time.Second * 3)
	fmt.Println("-----------")

	/*
		a := []int{1,2,3,4,5,6,7,8}
		for _, d := range a {
			fmt.Println(d,"--")
			if d==2{
				break
			}
			fmt.Println("+++")
		}


		runtime.GOMAXPROCS(1)
		intc := make(chan  int,1)
		stringc := make(chan string,1)
		intc <- 1
		stringc <- "fd"
		select {
			case va := <- intc:{
				fmt.Println(va)
			}
		case va := <- stringc:{
			panic(va)
		}

		}*/

}

/*判断是否是环形链表*/
func isRingList(root *domain.Node) bool {
	if root == nil {
		return false
	}
	preNode := root.Next.Next
	for root != nil && root.Next != nil && preNode != nil {
		if root == preNode {
			return true
		}
		root = root.Next
		preNode = preNode.Next.Next
	}
	return false
}

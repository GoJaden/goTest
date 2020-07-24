package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	userIds := []int{11, 13, 43, 43, 4324, 43242, 543, 65, 675}
	wg := sync.WaitGroup{}
	wg.Add(len(userIds))
	result := make(chan int, len(userIds)*2)
	for _, userId := range userIds {
		go func(userId int) {
			wg2 := sync.WaitGroup{}
			wg2.Add(2)
			go func(uId int) {
				fmt.Println("A正在处理", uId)
				time.Sleep(time.Second)
				result <- uId
				wg2.Done()
			}(userId)

			go func(uId int) {
				fmt.Println("B正在处理", uId)
				time.Sleep(time.Second)
				result <- uId
				wg2.Done()
			}(userId)
			wg2.Wait()
			wg.Done()
		}(userId)
	}
	wg.Wait()
	for i := 0; i < len(userIds)*2; i++ {
		select {
		case data := <-result:
			fmt.Println("处理完成的数据时", data)
		}
	}

}

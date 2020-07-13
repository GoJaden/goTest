package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.NewTicker(time.Second * 5)
	i := 0
	for {

		select {
		case d := <-t.C:
			{
				fmt.Println(strconv.Itoa(1), d)
				go func() {
					time.Sleep(time.Second * 1)
					i++
					fmt.Println("执行完毕任务", i)
				}()
			}

		}
	}
	/*aa := "1"
	fmt.Println(aa[:len(aa)-1])

	for i := 0; i < 100; i++ {
		go func(i int) {


		}(i)
	}
	time.Sleep(time.Second * 10)
	//var emptyStringIds []string = make([]string, 0)

	fmt.Println(IsGreaterThanGiveTime(1577243476, 2))*/
	/*fmt.Println(emptyStringIds)

	fmt.Println(emptyStringIds == nil)*/
	/*	var t *time.Timer

			t = time.NewTimer(time.Second*1)
		select {
		case <-t.C:
			fmt.Println("123")
			case <-t.C:
				fmt.Println("2222222")
		}*/

}

//判断某个时间是否是指定天数后的时间
func IsGreaterThanGiveTime(createTime int64, days int) bool {
	return time.Now().Unix() > time.Unix(createTime, 0).Add(time.Hour*24*2).Unix()
}

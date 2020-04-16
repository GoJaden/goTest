package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	aa := "1"
	fmt.Println(aa[:len(aa)-1])

	for i := 0; i < 100; i++ {
		go func(i int) {
			t := time.NewTimer(time.Second * 5)
			select {
			case <-t.C:
				fmt.Println(strconv.Itoa(i) + "----")

			}

		}(i)
	}
	time.Sleep(time.Second * 10)
	//var emptyStringIds []string = make([]string, 0)

	fmt.Println(IsGreaterThanGiveTime(1577243476, 2))
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

package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i) // 打印无规律数字
		}(i)
	}
	time.Sleep(time.Second)
}

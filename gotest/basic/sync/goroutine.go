package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	userCount := math.MaxInt64
	for i := 0; i < userCount; i++ {
		go func(i int) {
			// 做一些各种各样的业务逻辑处理
			fmt.Printf("go func: %d\n", i)
			time.Sleep(time.Second)
		}(i)
	}

	/*userCount := math.MaxInt64
	ch := make(chan bool, 2)
	for i := 0; i < userCount; i++ {
		ch <- true
		go Read(ch, i)
	}*/

	//time.Sleep(time.Second)
}

func Read(ch chan bool, i int) {
	fmt.Printf("go func: %d\n", i)
	<-ch
}

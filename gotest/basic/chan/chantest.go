package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	testChan := make(chan string)
	//testChan <- "test"
	go func() {
		for i := 0; i < 100; i++ {
			testChan <- "test" + strconv.Itoa(i)

		}
	}()

	/*go func() {
		for i:=100;i>0;i--{
			testChan <- "test"+strconv.Itoa(i)

		}
	}()*/
	/*timer := time.NewTimer(time.Second*10)
	select {
	case <-timer.C:{
		fmt.Println("同步阻塞3s")
	}
	}*/

	fmt.Println("---------")

	for i := 0; i < 200; i++ {
		a := <-testChan
		time.Sleep(time.Second * 3)
		fmt.Println(a)
	}

}

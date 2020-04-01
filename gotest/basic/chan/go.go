package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("+++")
	}()

	fmt.Println("---")

	time.Sleep(time.Second * 5)
	fmt.Println("~~~")
}

package main

import "fmt"

func main() {
	intFactory := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intFactory <- i
	}

	for {
		j := <-intFactory
		fmt.Println(j)
	}

}

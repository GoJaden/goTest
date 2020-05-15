package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {

		if i == 20 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("结束了")
}

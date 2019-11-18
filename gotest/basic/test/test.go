package main

import (
	"fmt"
	"time"
)

func printAll(vals []interface{}) { //1
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {

	/*names := []string{"stanley", "david", "oscar"}
	printAll(names)*/

	fmt.Println("hello")
	time.Sleep(time.Second * 10)
}

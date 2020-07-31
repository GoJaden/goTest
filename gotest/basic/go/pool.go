package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return 123
		},
	}

	l := myPool.Get() // [1]  调用 New
	fmt.Println(l)
	instance := myPool.Get() // [1] 调用 New
	fmt.Println(instance)
	myPool.Put(instance) // [2]
	l2 := myPool.Get()   // [3]
	fmt.Println(l2)
}

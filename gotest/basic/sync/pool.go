package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	pool := new(sync.Pool)

	for i := 0; i <= 1000; i++ {
		pool.Put("i=" + strconv.Itoa(i))
	}

	fmt.Println(pool.Get())

	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}

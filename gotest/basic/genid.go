package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	id int64
)

func genId() int64 {

	return atomic.AddInt64(&id, 1)
}

func main() {

	for i := 0; i < 10; i++ {
		go genId()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(id)
}

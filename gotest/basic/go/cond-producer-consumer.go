package main

import (
	"fmt"
	"sync"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	list := make([]interface{}, 0, 10)

	consumer := func() {
		cond.Signal()

		fmt.Println("remove")
		list = list[1:]

	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		if len(list) == 1 {
			cond.Wait()
		}
		fmt.Println("add element...")
		list = append(list, struct{}{})
		go consumer()
		cond.L.Unlock()
	}

}

func producer() {

}

func consumer() {

}

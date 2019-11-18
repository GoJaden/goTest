package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Id    int
	Mutex sync.Mutex
}

func main() {
	per := Person{
		Id:    111,
		Mutex: sync.Mutex{},
	}
	per.Mutex.Lock()
	fmt.Println(per)
	per.Mutex.Unlock()

}

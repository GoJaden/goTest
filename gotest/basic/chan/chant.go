package main

import (
	"fmt"
	"unsafe"
)

func main() {

	inr := unsafe.Pointer(new(interface{}))
	fmt.Println(inr)
	ch := make(chan int, 0)

	close(ch)

	dat, ok := <-ch
	if ok {
		fmt.Println("--", dat)
	}
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

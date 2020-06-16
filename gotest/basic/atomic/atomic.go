package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var old int32 = 23
	var addr *int32 = &old
	new := 100
	b := atomic.CompareAndSwapInt32(addr, old, int32(new))
	fmt.Println(b)

}

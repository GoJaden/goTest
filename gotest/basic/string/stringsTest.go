package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var aa uint64 = 123213211232131
	var cc int = int(aa)
	fmt.Println(cc)
	bb := strconv.FormatUint(aa, 10)
	fmt.Println(bb)

	t := time.Now()
	fmt.Println(t)
}
func t() {
	fmt.Println("123")
}

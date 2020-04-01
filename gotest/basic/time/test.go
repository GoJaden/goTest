package main

import (
	"fmt"
	"time"
)

func main() {
	D := time.Now().UnixNano()
	time.Sleep(time.Second*2)
	DA := (time.Now().UnixNano() - D)/ 1000000
	fmt.Println(DA)
}

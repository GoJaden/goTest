package main

import (
	"crypto/rand"
	"fmt"
)

func main() {

	data, _ := rand.Read([]byte{0, 0, 0, 0, 0, 0})
	fmt.Println(data)
}

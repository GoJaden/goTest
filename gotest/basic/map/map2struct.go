package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	bts, _ := cmd.CombinedOutput()
	fmt.Println(string(bts))
	MapToStruct()

	cc := make(chan int, 5)
	cc <- 1
	close(cc)
	fmt.Println(<-cc)
}

func MapToStruct() {

	aMap := make(map[interface{}]interface{})
	aMap["string"] = 123
	aMap[123] = "3213"

	for k, da := range aMap {
		fmt.Println(k, ":", da)
	}
}

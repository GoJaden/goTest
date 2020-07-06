package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type A struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
	Info string `json:"info,omitempty"`
}

func main() {

	fmt.Println(math.MaxInt64 > 4102415940)

	aa := &A{
		Age:  0,
		Info: "",
	}

	bts, _ := json.Marshal(aa)
	fmt.Println(string(bts))

	a := []byte{'1', '2', '3'}

	fmt.Println(string(a))

	var b []byte
	b = nil
	fmt.Println(b)

	fmt.Println(string(b))
}

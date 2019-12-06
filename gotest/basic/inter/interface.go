package main

import (
	"fmt"
	"reflect"
)

func main() {
	b := reflect.TypeOf(getTest())
	fmt.Println(b)

}

func getTest() interface{} {
	return "12"
}

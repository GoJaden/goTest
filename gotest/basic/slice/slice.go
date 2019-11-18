package main

import (
	"fmt"
	"reflect"
)

func main() {

	slice := new([]string)
	fmt.Println(slice)
	fmt.Println(reflect.ValueOf(slice).Kind())

}

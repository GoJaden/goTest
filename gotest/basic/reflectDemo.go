package main

import (
	"fmt"
	"reflect"
)

func main() {

	var test int = 0

	fmt.Println(reflect.TypeOf(test))

	v := reflect.ValueOf(test)

	fmt.Println(v.Kind().String())

	if v.CanSet() {
		v.SetInt(10)
	}

	fmt.Println(v, "-", test)

}

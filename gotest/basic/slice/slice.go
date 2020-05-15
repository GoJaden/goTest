package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	result := strings.Split("app-active bbb  sf", " ")
	if len(result) > 0 {
		fmt.Println(result, len(result))
	}

	aa := []int{1, 32, 42, 12, 13, 45, 67, 6, 7, 3, 55, 21, 43, 53, 3, 43, 4343}
	sort.Slice(aa, func(i, j int) bool {
		if aa[i] > aa[j] {
			return true
		}
		return false
	})
	fmt.Println(aa)

	slice := new([]string)
	fmt.Println(slice)
	fmt.Println(reflect.ValueOf(slice).Kind())

	json.Marshal(slice)
}

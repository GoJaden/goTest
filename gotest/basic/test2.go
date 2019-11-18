package main

import (
	"common-lib/collecttion"
	"fmt"
)

func main() {

	/*var buffer []byte=make([]byte,1024)
	var buffer2 []byte

	fmt.Println(len(buffer))
	fmt.Println(len(buffer2))*/

	/*result := time.Duration(1000000)//+time.Second

	fmt.Println(result)*/

	set := collecttion.NewHashSet()
	set.Add(123)
	var myMap map[string]int = make(map[string]int, 5)
	myMap["namg"] = 1
	myMap["age"] = 23
	myMap["fdf"] = 3223

	delete(myMap, "age")

	myMap["fdf"] = 110
	fmt.Println(myMap["fdf"])
	fmt.Println(set.Exist(123))
}

type HashSet struct {
	data map[interface{}]struct{}
}

func (h *HashSet) Add(value interface{}) {
	h.data[value] = struct{}{}
}

package main

import (
	"fmt"
	"strings"
)

func printAll(vals []interface{}) { //1
	for _, val := range vals {
		fmt.Println(val)
	}
}

type S struct {
	A string
}

func main() {

	go func() {
		fmt.Print("t1")
		go func() {
			fmt.Print("t2")
			go func() {
				fmt.Print("t3")
			}()
		}()
	}()

	var s float32 = 3.1

	fmt.Print(s)

	str := "abc" + "123"
	fmt.Println(str)

	fmt.Println(ParseExaminers("2"))
}
func ParseExaminers(examiners string) []string {
	if strings.Contains(examiners, ",") {
		return strings.Split(examiners, ",")
	}
	return []string{examiners}
}

type T struct {
}

func (t T) test() {

}

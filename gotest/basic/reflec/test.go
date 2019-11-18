package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode"
)

type Person struct {
	Id    *int
	Name  string
	Level *int
}

func main() {
	/*var a int = 1
	p := Person{
		Id:    nil,
		Name:  "123",
		Level: &a,
	}
	if p.Id == nil{

	}
	reflect.ValueOf(p.Id).Elem().SetInt(100)
	fmt.Println(reflect.ValueOf(p.Id).Elem().Int())*/

	var a *A
	fmt.Println(reflect.ValueOf(a), reflect.TypeOf(a))

	fmt.Println(parseSuffixNumber("qqqqqew12ds31231"))
}

type A struct {
}

//获取问题标记的后缀
func parseSuffixNumber(str string) int {
	bs := []byte(str)
	for poi := len(bs) - 1; poi >= 0; poi-- {
		if unicode.IsLetter(rune(bs[poi])) {
			if len(bs) != poi+1 {
				fmt.Println(string(bs[poi+1:]))
				number, err := strconv.Atoi(string(bs[poi+1:]))
				if err != nil {
					fmt.Println("问题的序号错误,"+str, err.Error())
					return 0
				}
				return number
			} else {
				return 0
			}
		}
		continue
	}
	return 0
}

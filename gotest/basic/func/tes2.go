package main

import (
	"fmt"
	"reflect"
)

//专门用于参数校验
type CheckParams func(args interface{}) Inter

type Inter interface {
	DoSomething()
}

type Person struct {
	Id int
}
type NewPerson struct {
	Id int
}

func (p NewPerson) DoSomething() {
	fmt.Println(p.Id, "新的处理方式正在处理业务...")
}

func (p Person) DoSomething() {
	fmt.Println(p.Id, "正在处理业务...")
}

func main() {
	var d CheckParams = check
	d(111).DoSomething()
	//fmt.Printf("%v,%T",,d)
}

func check(args interface{}) Inter {
	var res int = 0
	switch reflect.TypeOf(args).Kind() {
	case reflect.Int:
		v := reflect.ValueOf(args)
		res = int(v.Int())

	default:
		fmt.Println("----------格式错误")
	}
	p := NewPerson{Id: res}
	return p
}

package main

import (
	"fmt"
	"reflect"
)

type Req struct {
	id    int
	name  string
	isVip bool
	time  int64
}

type User struct {
	Id       int
	Name     string
	Title    string
	Birthday int64
	Info     Info
}
type Info struct {
	Id  int
	SSS string
}

type Func interface {
	do()
}

func (r Req) do() {
	fmt.Println("do...")
}

func StructToString(st interface{}) string {
	k := reflect.TypeOf(st).Kind()
	if k == reflect.Struct {

		return "123"
	} else if k == reflect.Ptr {
		fmt.Println(reflect.ValueOf(st).Type())
	}
	return "1"
}

func main() {

	fmt.Println(StructToString(&Info{}))
	/*dst :=(*Func)(nil)
	dstType := reflect.TypeOf(dst).Elem()
	var re Req =Req{}
	s := reflect.TypeOf(re).Implements(dstType)
	fmt.Println(s)

	var aa interface{}="132"
	var bb string="132"
	v1 :=reflect.ValueOf(bb)
	 c :=v1.Interface()*/

	//var t reflect.Type=reflect.TypeOf(aa)
	//var tk reflect.Kind=t.Kind()
	//var v reflect.Value = reflect.ValueOf(aa)
	//fmt.Println(c,v)
}

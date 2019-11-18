package main

import (
	"fmt"
	"reflect"
)

type Te struct {
	Id   int
	Name string
	Age  uint8
	Vip  bool
	In   Inner
}

type Inner struct {
	Id         int
	InngerName string
}

var tes struct {
}

type Func2 interface {
	do2()
}

func (t Te) do2() {
	fmt.Println("do........")
}

func main() {
	t := Te{1, "jaden", 12, false, Inner{
		Id:         32132,
		InngerName: "test",
	}}
	/*ty := reflect.TypeOf(&t)

	fmt.Println(ty)*/

	getObjectInfo(t)
}

/**
获取结构体对象的信息
*/
func getObjectInfo(o interface{}) {
	ty := reflect.TypeOf(o)
	if k := ty.Kind(); k != reflect.Struct {
		fmt.Println("必须获取结构体对象的信息...")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println(ty.Name(), v)

	for i := 0; i < v.NumField(); i++ {
		//f := ty.Field(i)
		val := v.Field(i).Interface()
		t1 := reflect.TypeOf(val)
		if k := t1.Kind(); k == reflect.Struct {
			getObjectInfo(val)
		}

	}

	for i := 0; i < ty.NumMethod(); i++ {
		m := ty.Method(i)
		fmt.Printf("%s:%v \n", m.Name, m.Type)
	}

}

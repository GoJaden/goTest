package main

import (
	jsonutil "common-lib/json-util"
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type Parent struct {

}
func (p Parent) Marsh(this interface{}){
	fmt.Println(jsonutil.ObjectToJsonString(this))
}

type ChildOne struct {
	Parent
	Name string
	Age int
	Info interface{}
}
//声明类型别名   MyType 跟 int是相同的类型
type MyType = int
// type MyType int    MyType 跟 int不是同一类型

func main() {
	var a MyType = 12

	var b int =a

	fmt.Println(b)
	go func() {

		for i := 0; i < 100; i++ {
			go func() {
				time.Sleep(time.Second)
			}()
		}
	}()
	var c int = 1.0
	fmt.Println(c,runtime.NumCPU(),runtime.NumGoroutine())
	o := ChildOne{
		Name:   "123",
		Age:    11,
		Info:   "123",
	}
	o.Marsh(o)
	o1 := ChildOne{
		Name:   "123",
		Age:    11,
		Info:"123",
	}
	fmt.Println(o==o1)
	bts ,err := json.Marshal(nil)

	fmt.Print(string(bts),err)

}

package main

import "fmt"

//对这个字符串处理
type Func func(name string) string

func main() {

	var a Func = dealName

	fmt.Println(a("jaden"))

}

//申明一个 参数 以及返回值类型一样的 函数，可以实现多态
func dealName(name string) string {
	return name + "fff"
}

package main

import (
	"fmt"
	"strings"
)

type Info struct {
	Page struct {
		Total    int `json:"total"`
		PageSize int `json:"pageSize"`
		PageNum  int `json:"pageNum"`
	} `json:"page"`
	List []*B `json:"list"`
}
type B struct {
	Id int
}

/**
结构体不管是new 出来的对象还是以结构体方式声明的对象，包括里层的对象，都是有分配空间的
new的是对象指针    结构体只是对象的值

对象内层的切片这种对象，是分配有  零  空间的，必须手动扩展空间，方可使用

golang对程序的优化
//info.List[0].Id=12  √
data := info.List[0]  ×
data.Id=12
*/

func main() {
	//s := new([]string)
	info := new(Info)
	fmt.Println("%T", info)
	info.List = make([]*B, 2)
	//info := new(Info)
	//info.List = append(info.List,B{Id:1})
	//info.List= make([]*B,2)
	//info.List[0].Id=12
	info.List[0] = new(B)
	fmt.Println("%T", info.List[0].Id)
	/*data.Id=12
	info.List[0]=data
	fmt.Println(info)*/

	s := make(map[string]int)
	fmt.Println(s)

	params := []interface{}{1}
	params = append(params, 2)

	p := strings.Repeat("?,", 10)
	fmt.Println(p)
}

func Handler(b *B) {

}

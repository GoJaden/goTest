package main

import (
	"fmt"
	"sort"
)

type MyType map[int]map[string]int

func (m MyType) Len() int {
	return len(m)
}

func (m MyType) Less(i, j int) bool {
	return i > j
}

func (m MyType) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	t := make(map[int]map[string]int)
	v := make(map[string]int)
	v["zhangsan"] = 111
	v["lisi"] = 333
	v["wangwu"] = 122
	if _, exist := v["name"]; exist {
		fmt.Println("已经存在")
	}

	t[0] = v
	t[11] = v
	t[2] = v
	if _, exist := t[2]; exist {
		delete(t, 2)
	}
	sort.Sort(MyType(t))

	fmt.Println(t)
}

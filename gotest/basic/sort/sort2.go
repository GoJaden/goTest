package main

import (
	"fmt"
	"sort"
)

func main() {

	s1 := P2{P1{1, "jj", 10}}
	s2 := P2{P1{2, "jj", 30}}
	s3 := P2{P1{3, "jj", 40}}
	s4 := P2{P1{4, "jj", 20}}
	s5 := P2{P1{5, "jj", 70}}
	s6 := P2{P1{6, "jj", 60}}
	s7 := P2{P1{7, "jj", 88}}
	res1 := make(q, 0)
	res1 = append(res1, s1)
	res1 = append(res1, s2)
	res1 = append(res1, s3)
	res1 = append(res1, s4)
	res1 = append(res1, s5)
	res1 = append(res1, s6)
	res1 = append(res1, s7)

	//todo

	sort.Sort(res1)

	fmt.Println(res1)

}

type p []P1

type P1 struct {
	Id   int
	Name string
	Age  int
}
type q []P2
type P2 struct {
	P1
}

func (u q) Len() int {
	return len(u)
}

func (u q) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}
func (u q) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u p) Len() int {
	return len(u)
}

func (u p) Less(i, j int) bool {
	return u[i].Id < u[j].Id
}
func (u p) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	users := make(Users, 0)
	u1 := User{
		Id:   0,
		Name: "zs",
	}
	u2 := User{
		Id:   222,
		Name: "z323s",
	}
	u3 := User{
		Id:   2,
		Name: "1z1s",
	}
	u4 := User{
		Id:   1,
		Name: "32zs",
	}
	u5 := User{
		Id:   3232,
		Name: "zdss",
	}
	users = append(users, u1)
	users = append(users, u2)
	users = append(users, u3)
	users = append(users, u4)
	users = append(users, u5)
	sort.Sort(users)
	for _, u := range users {
		fmt.Println(u)
	}
}

type Users []User
type User struct {
	Id   int
	Name string
}

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Id < u[j].Id
}
func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

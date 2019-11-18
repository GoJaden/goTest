package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	*Info `json:"1"`
}
type Info struct {
	Id       int    `json:"id"`
	InfoTest string `json:"info_test"`
}

func main() {

	s := Stu{
		Id:   1,
		Name: "jaden",
	}

	da, _ := json.Marshal(s)
	fmt.Println(string(da))

	ss := "{\"id\":1,\"name\":\"jaden\"}}"
	sty := new(Stu)
	json.Unmarshal([]byte(ss), sty)
	fmt.Println(sty)
}

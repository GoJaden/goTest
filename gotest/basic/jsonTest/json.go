package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var test int = 55
	q, _ := json.Marshal(test)
	fmt.Println(string(q))

	a := new(A)
	a.Id = 1
	a.Name = "jaden"
	bb := make([]B, 0)
	bb = append(bb, B{
		Id:   12,
		Info: "fdf",
	}, B{
		Id:   13,
		Info: "323",
	},
		B{
			Id:   14,
			Info: "321",
		})
	a.Bs = bb
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
	aaa := new(A)
	json.Unmarshal([]byte("{\"id\":1,\"name\":\"jaden\",\"Bs\":[{\"_id\":12,\"info\":\"fdf\"},{\"_id\":13,\"info\":\"323\"},{\"_id\":14,\"info\":\"321\"}]}"), aaa)
	fmt.Println(aaa)
}

type A struct {
	Id   interface{} `json:"id"`
	Name interface{} `json:"name"`
	Bs   []B
}

type B struct {
	Id   int    `json:"_id"`
	Info string `json:"info"`
}

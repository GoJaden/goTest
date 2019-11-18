package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"goProject/gotest/basic/protobuftest"
	"io/ioutil"
	"os"
)

func main() {
	var id uint32 = 110
	idptr := &id

	name := "jaden"
	nameptr := &name

	isvip := false
	isvipptr := &isvip

	p := &protobuftest.Person{
		Id:    idptr,
		Name:  nameptr,
		IsVip: isvipptr,
	}
	buffer := make([]byte, 1024)
	buffer, err1 := proto.Marshal(p)
	if err1 != nil {
		fmt.Println(err1)
	}

	/*err :=p.XXX_Unmarshal(buffer)
	if err != nil{
		fmt.Println("出错了",err)
	}*/
	e := ioutil.WriteFile("user.txt", buffer, os.ModeAppend)
	if e != nil {
		fmt.Println(e)
	}
	iobuffer := make([]byte, 1024)
	iobuffer, err2 := ioutil.ReadFile("user.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	p2 := &protobuftest.Person{}
	err3 := proto.Unmarshal(iobuffer, p2)

	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(p2)
}

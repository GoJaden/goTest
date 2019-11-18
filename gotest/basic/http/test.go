package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://doc.maizuo.com/api/file/getAttach?fileId=5b04cb63915399000e000008")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buffer))
}

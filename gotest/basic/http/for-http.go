package main

import (
	httputil "common-lib/http-util"
	"fmt"
	"time"
)

func main() {
	go func() {
		for i:=0;i<100;i++{
			_,_,err := httputil.NewHttpClient().DoGet("http://www.baidu.com","").Response()
			if err != nil{
				fmt.Println("err:",err)
			}
			fmt.Println(i)
		}
	}()

	fmt.Println("--")
	time.Sleep(time.Second*20)
}

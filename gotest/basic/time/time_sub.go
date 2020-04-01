package main

import (
	httputil "common-lib/http-util"
	"fmt"
	"time"
)





func main() {
	fmt.Println(IsGreaterThanGiveTime(time.Now().Add(-time.Hour*24*2).Unix(),2))
	return




	resp ,result,err := httputil.NewHttpClient().SetTimeout(time.Nanosecond*1).DoGet("http://www.baidu.com","").Response()
	if err != nil{
		if err.Error() =="Get http://www.baidu.com: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"{
			fmt.Println("---------")
			return
		}
		fmt.Println(err)
	}
	fmt.Println(resp,result)
	beforeFiveMinuteTime := time.Now().Add(-time.Minute*5).Local().Format("2006-01-02 15:04:05")

	fmt.Println(beforeFiveMinuteTime)
	var a = 11
	ab(&a)
}

func ab(a *int){
fmt.Println(a)
}
package main

import (
	"fmt"
	"time"
)

func main() {
	/*now := time.Now().Format("2006-01-02 03:04:05")
	fmt.Println(now)
	t1 ,_:=time.ParseInLocation("2006-01-02 15:04:05",time.Now().Format("2006-01-02 15:04:05"),time.Local)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(t1)
	ti :=time.Unix(1553673548,0).Unix()
	fmt.Println(ti)*/

	now := time.Now().Unix()
	/*tt,_:=time.Unix(now,0).MarshalJSON()
	fmt.Println("json",string(tt))

	tx ,_:= time.Unix(now,0).MarshalText()

	fmt.Println(string(tx[0:10]))*/
	fmt.Println(FormatToDate(time.Unix(now, 0)))
}
func FormatToDate(timeObj time.Time) string {
	return timeObj.Format("2006-01-02")
}

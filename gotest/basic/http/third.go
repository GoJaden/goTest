package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	const (
		MEINIAN_SUPPLY  = iota + 1 //美年
		KANGKANG_SUPPLY            //康康
	)
	fmt.Println(MEINIAN_SUPPLY, KANGKANG_SUPPLY)

	/*i :=0
	for i<1024{
		i++
		fmt.Println(string(i))
	}
	var B string="123"
	var V string="adfdf"
	s := fmt.Sprintf(B+"/qwer"+V+"/%s,%s","qqq","aaa")
	fmt.Println(s)*/
	/*s := EncryptToMD5([]byte("GetAllCity201910102310095204pL6iqun4McX9Nramd36dc9fd-3de9-476c-aeb6-c22ba3198547"))
	url := fmt.Sprintf("http://m.51kys.cn:6677/open/openapi/PackageService.ashx?action=GetAllCity&onlyCode=d36dc9fd-3de9-476c-aeb6-c22ba3198547&appId=201910102310095204&signature=%s",s)
	resp ,s ,err := httputil.NewHttpClient().DoGet(url,"").Response()
	fmt.Println(resp,s,err)*/
}
func EncryptToMD5(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

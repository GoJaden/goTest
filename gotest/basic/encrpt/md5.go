package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	/*
		查询订单根据订单Id
		查询订单根据订单Id，返回订单状态信息！
		请求URL
		http://m.kktijian.cn:6677/openapi/OrederService.ashx?action=serachOrderByorderId&onlyCode=47146CFF163F4EE0BECF076A3A960527&appId=12345678&orderId=27196&signature=ab25596d2e3e08ed5d4cc41b77c80f2b
		传入参数说明
		 action：serachOrderByorderId，此参数为固定请求方法参数，查询订单
		 onlyCode：分销商唯一码，如：47146CFF163F4EE0BECF076A3A960527；
		 appId：应用id；如：12345678；
		 appKey：应用key；如：kyi9GyLs；
		 orderId：订单编号，如：27196，
		签名规则
		 签名所传入的参数与字符串组合顺序必须与示例相同：
		 string signText = action + appKey + appId + onlyCode + orderId;
		 signature=MD5(signText)；*/

	/*"onlyCode": "309adcad-a468-45ba-be0f-be0a8f6ffaad",
		"appId": "201910311553545734",
		"appKey": "sbazi0iI4F00Wes3",
		"supplyId": 323,
		"prefix": "kk_"

	fmt.Sprint(r.credential.AppKey, r.credential.AppId, r.OnlyCode, r.OrderId)                    2019110517141600000027a5779de4f8*/
	a := "cancelOrderByorderIdsbazi0iI4F00Wes3201910311553545734309adcad-a468-45ba-be0f-be0a8f6ffaad2019110517141600000027a5779de4f8"
	//a := "serachOrderByorderIdsbazi0iI4F00Wes3201910311553545734309adcad-a468-45ba-be0f-be0a8f6ffaad2019110517141600000027a5779de4f8"

	//http://api.kktijian.com/openapi/OrederService.ashx?action=serachOrderByorderId&onlyCode=309adcad-a468-45ba-be0f-be0a8f6ffaad&appKey=sbazi0iI4F00Wes3&appId=201910311553545734&orderId=2019110517141600000027a5779de4f8&signature=81a795b0cfc77764a234abbab9832ab5
	fmt.Println("ab222e08958522d3e5a7d9073a1e4565" == "ab222e08958522d3e5a7d9073a1e4565")
	fmt.Println(EncryptToMD5([]byte(a)))
}

func EncryptToMD5(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

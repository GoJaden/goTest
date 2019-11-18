package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	params := make(map[string]string)
	params["A"] = "疯掉了健康"
	params["B"] = "12"
	params["C"] = "挂法规的防辐射服"
	params["D"] = "佛挡杀佛费点事"

	paramStr := ""
	for k, v := range params {
		paramStr = paramStr + fmt.Sprintf("&%s=%s", k, url.QueryEscape(v))
	}
	fmt.Println(paramStr)
	fmt.Println(strings.Index(paramStr, "%E"))

	/*url := url.QueryEscape("signType=1&dateStr=20191104101237&A=752&B=13163251830&C=2019110409532195215&D=白领体检套餐A（女未婚）&E=2019-11-06&F=二七分院&G=周二--周日07;30——17:00 周一：07:30-12：00&H=南三环嵩山南路东南角，美年大健康体检中心（艾菲尔时尚酒店隔壁）")
	fmt.Println(url)*/
	//url.decode
	/*q ,_:=url.QueryUnescape("%E4%B8%AD%E5%9B%BD")
	//fmt.Println(url)
	fmt.Println(q)

	u := "http://www.baidu.com?query=123&data=123"
	bo ,_:=url.Parse(u)
	fmt.Println(bo.Host)

	us := url.Values{}
	us.Add("data","123")
	us.Add("data","456")
	us.Add("data","789")
	us.Add("data","012")
	us.Add("par","ad")
	fmt.Println(us.Encode())

	fmt.Println("--------------------")
	fmt.Println("++",0010)
	rand.Seed(100)
	for i:=0;i<100;i++{
		fmt.Println(rand.Intn(10))
	}*/

}

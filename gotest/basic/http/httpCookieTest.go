package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type UserDo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	User     string `json:"user"`
}

func main() {

	/*resp ,err := http.Get("http://www.baidu.com")
	if err != nil{
		fmt.Println("get失败",err)
		return
	}
	defer resp.Body.Close()
	cks := resp.Cookies()
	for _,val := range cks{
		fmt.Println(val)
	}*/

	/*body := &UserDo{
		Email:    "",
		Password: "jaden123",
		User:     "jaden",
	}
	bts ,_ := json.Marshal(body)
	r := strings.NewReader(string(bts))

	r.Read(bts)
	fmt.Println(string(bts))
	client := http.Client{
	}
	resp , err := client.Post("http://grafana.maizuo.com/login","application/json",r)
	if err != nil{
		fmt.Println("errr",err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("status:",resp.Status)
	cks := resp.Cookies()
	for _,val := range cks{
		fmt.Println(val)
	}*/

	url := "http://grafana.maizuo.com/login"
	method := "POST"

	payload := strings.NewReader("{\"user\":\"jaden\",\"email\":\"\",\"password\":\"jaden123\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Cookie", "grafana_session=c10063bf95f57d561a893e966a812b3b")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	cks := res.Cookies()
	for _, val := range cks {
		fmt.Println(val)
	}
}

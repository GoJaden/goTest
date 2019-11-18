package main

import (
	jsonutil "common-lib/json-util"
	"fmt"
	"net/url"
	"sdyxmall/business-examine/utils"
	"time"
)

func main() {
	/*"onlyCode": "309adcad-a468-45ba-be0f-be0a8f6ffaad",
		"appId": "201910311553545734",
		"appKey": "sbazi0iI4F00Wes3",
		"supplyId": 323,
		"prefix": "kk_"

	fmt.Sprint(r.credential.AppKey, r.credential.AppId, r.OnlyCode, r.OrderId)*/
	//取消预约
	/* req := ReqExamineCancel{
	           Appkey:   "sbazi0iI4F00Wes3",
	           AppId:    "201910311553545734",
	           OnlyCode: "309adcad-a468-45ba-be0f-be0a8f6ffaad",
	           OrderId:  "857420",
	   }
	   fmt.Println("/openapi/OrederService.ashx?"+makeQueryString("cancelOrderByorderId",&req))*/

	t, err := time.Parse("2006-01-02 15:04:05", "2019-11-07 15:10:01")
	fmt.Println(t)
	if err != nil {
		fmt.Println(err)
	}
	this := t.Unix() + (3*60*60 - 5411)
	fmt.Println(this)

	fmt.Println(time.Unix(this, 0).In(time.UTC).String())
	req := &ReqExamineOrderInfo{
		Appkey:   "sbazi0iI4F00Wes3",
		AppId:    "201910311553545734",
		OnlyCode: "309adcad-a468-45ba-be0f-be0a8f6ffaad",
		OrderId:  "859448",
	}
	fmt.Println("http://api.kktijian.com/openapi/OrederService.ashx?" + makeQueryString("serachOrderByorderId", req))

}

type ReqExamineOrderInfo struct {
	Appkey   string `json:"-"`
	AppId    string `json:"appId"`
	OnlyCode string `json:"onlyCode"`
	OrderId  string `json:"orderId"`
}

func (s ReqExamineOrderInfo) Source() string {
	return fmt.Sprint(s.Appkey, s.AppId, s.OnlyCode, s.OrderId)
}

type Request interface {
	Source() string
}
type ReqExamineCancel struct {
	Appkey   string `json:"-"`
	AppId    string `json:"appId"`
	OnlyCode string `json:"onlyCode"`
	//订单编号，如：27196，
	OrderId string `json:"orderId"`
}

func (r *ReqExamineCancel) Source() string {
	//action + appKey + appId + onlyCode + orderId;
	return fmt.Sprint(r.Appkey, r.AppId, r.OnlyCode, r.OrderId)
}
func makeQueryString(action string, req Request) string {
	urlValues := toUrlValues(req)
	urlValues.Add("action", action)
	urlValues.Set("signature", getSignature(action+req.Source()))
	return urlValues.Encode()
}

func toUrlValues(data interface{}) *url.Values {
	valueMap := map[string]string{}
	if e := jsonutil.ByteToObject(jsonutil.ObjectToJsonBytes(data), &valueMap); e != nil {
		panic(e)
	}
	values := &url.Values{}
	for key, value := range valueMap {
		values.Add(key, value)
	}
	return values
}

func getSignature(args ...interface{}) string {
	s := fmt.Sprint(args...)
	return utils.EncryptToMD5([]byte(s))
}

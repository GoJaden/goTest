package main

import (
	httputil "common-lib/http-util"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sdyxmall/business-examine/utils"
	"strconv"
	"time"
)

//MTdhZTYzMjUxOGEzYTQ2YTJjYzNiYjhhMTliZTExMTE=

func main() {
	req := ReqExamine{
		ExaminePid: "TJW52004132215174343",
		UId:        "12",
		Token:      "ZmViNTg3ZjczYTlhYzQzNDYyMDA5NjQzYjFlZGNiZGY=",
	}
	s, err := md5Signer.SignUrlWithQuery("https://api.meinian365.com//Medical/ExamineOrderInfo", req)
	if err != nil {
		fmt.Println(err)
	}
	resp, result, err := httputil.NewHttpClient().DoGet(s, "").Response()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	fmt.Println(result)
}

var md5Signer = Signer{}

type Signer struct{}

type ReqExamine struct {
	ExaminePid string `json:"examine_pid"`
	UId        string `json:"uid"`
	Token      string `json:"token"`
}

// 对queryString签名 然后与rawURL拼接返回
func (s Signer) SignUrlWithQuery(rawUrl string, queryParameter interface{}) (string, error) {
	targetURL, e := url.Parse(rawUrl)
	if e != nil {
		return "", e
	}
	queryStringWithSignature, e := s.getSignedUrlValues(queryParameter)
	if e != nil {
		return "", e
	}
	targetURL.RawQuery = queryStringWithSignature.Encode()
	return targetURL.String(), nil
}

const (
	_fieldEncryptSign = "encryptsign"
	_fieldTimestamp   = "timestamp"
)

// 对结构体签名并返回map对象
func (s Signer) SignBodyData(data interface{}, withTimestamp bool) (map[string]string, error) {
	if data == nil {
		return nil, errors.New("input cannot be null")
	}
	paramMap, e := s.parseQueryStruct(data)
	if e != nil {
		return nil, e
	}
	if withTimestamp {
		paramMap[_fieldTimestamp] = strconv.Itoa(int(time.Now().Unix()))
	}
	signature, e := signJSONBytes(paramMap)
	if e != nil {
		return nil, e
	}
	paramMap[_fieldEncryptSign] = signature
	return paramMap, nil
}

// 对传入参数签名并生成url.Values
func (s Signer) getSignedUrlValues(data interface{}) (*url.Values, error) {
	dataMap, e := s.SignBodyData(data, true)
	if e != nil {
		return nil, e
	}
	urlValues := &url.Values{}
	for paramName, paramValue := range dataMap {
		urlValues.Set(paramName, paramValue)
	}
	return urlValues, nil
}

// 对数据进行JSON序列化后MD5
func signJSONBytes(data interface{}) (string, error) {
	dataJSONBytes, e := json.Marshal(data)
	if e != nil {
		return "", e
	}
	return utils.EncryptToMD5(dataJSONBytes), nil
}

func (s Signer) parseQueryStruct(content interface{}) (result map[string]string, e error) {
	if marshalContent, err := json.Marshal(content); err != nil {
		e = err
		return
	} else {
		var val map[string]interface{}
		result = make(map[string]string)
		if err := json.Unmarshal(marshalContent, &val); err != nil {
			e = err
			return
		} else {
			for k, v := range val {
				var queryVal string
				switch t := v.(type) {
				case string:
					queryVal = t
				case float64:
					queryVal = strconv.FormatFloat(t, 'f', -1, 64)
				case time.Time:
					queryVal = t.Format(time.RFC3339)
				default:
					j, err := json.Marshal(v)
					if err != nil {
						continue
					}
					queryVal = string(j)
				}
				result[k] = queryVal
			}
		}
	}
	return
}

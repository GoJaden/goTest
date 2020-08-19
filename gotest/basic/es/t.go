package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type ScheduleReq struct {
	bid      int
	CinemaId int
	//格式为yyyyMMdd
	ShowDate string
	//请求时间戳
	dateTime string
	sign     string
}

func main() {
	/*tm := time.Unix(time.Now().Unix(), 0)
	s := tm.Format("20060102150405")

	ss := fmt.Sprintf("%s",s)

	req := &ScheduleReq{
		bid:      12,
		CinemaId:323,
		ShowDate: "2020-08-17",
		dateTime: ss,
		sign:     "11",
	}
	url := fmt.Sprintf("/seat/foretells.htm?bid=%v&cinemaid=%v&showdate=%s&datetime=%s&sign=%s",
		req.bid, req.CinemaId, req.ShowDate, req.dateTime, req.sign)
	fmt.Println(url)
	*/
	/*
		http://api.kktijian.com/openapi/PackageService.ashx?action=GetPackageBypackageCode&onlyCode=309adcad-a468-45ba-be0f-be0a8f6ffaad&appId=201910311553545734&packageCode=TC201907171010420001&signature=51045c4c85a6d3d03dc370ab1c75e143
	*/
	/*
	   string signText = action + appId + appKey + onlyCode + packageCode;签名字符串
	    signature=MD5(signText)；
	*/
	tt := "getModelByHospId201910311553545734sbazi0iI4F00Wes3309adcad-a468-45ba-be0f-be0a8f6ffaadYY201709141548200001"

	fmt.Println(EncryptToMD5([]byte(tt)))
}

func EncryptToMD5(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

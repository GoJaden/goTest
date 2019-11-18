package main

import (
	timeutil "common-lib/time-util"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(IsEqualParseIdCardBirthday("421302199602148416", 824227200))

}

//校验身份证信息 与 填写的生日是否 一致
func IsEqualParseIdCardBirthday(idcard string, birthday int64) bool {
	if len(idcard) < 15 || birthday == 0 {
		return false
	}
	splitBirthday := idcard[6:14]
	fmt.Println(splitBirthday)
	bt := timeutil.Parse("20060102", splitBirthday)
	if bt == nil {
		return false
	}
	if bt.Unix() == birthday {
		return true
	}
	return false
}

//校验身份证信息 与 填写的性别是否一致
func IsEqualParseIdCardSex(idcard string, sex int) bool {
	if len(idcard) < 15 || sex <= 0 {
		return false
	}
	parseSex := idcard[len(idcard)-2 : len(idcard)-1]
	fmt.Println(parseSex)
	parseSexNumber, err := strconv.Atoi(parseSex)
	if err != nil {
		return false
	}
	if parseSexNumber%2 == 0 && sex == 2 {
		return true
	}
	if parseSexNumber%2 != 0 && sex == 1 {
		return true
	}
	return false
}

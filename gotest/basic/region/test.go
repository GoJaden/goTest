package main

import (
	"errors"
	"fmt"
)

func main() {
	/*c ,err:=GetDistrictLevel("440012")
	fmt.Println(c,err)*/
	qq := `123455`
	qq += ` 123111`
	qq += ` qwe`
	fmt.Println(qq)

}

func GetDistrictLevel(code string) (int, error) {
	if len(code) == 0 || len(code) != 6 {
		return 0, errors.New(code + "区域编码有误")
	}
	if code[4] == '0' && code[5] == '0' {
		if code[2] == '0' && code[3] == '0' {
			return 0, nil
		}
		return 1, nil
	}
	return 2, nil
}

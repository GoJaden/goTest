package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

/***
sql中包括常见查询有in 、not in、like、group by、order by、between and 、
*/

func main() {
	regs := make([]interface{}, 0)
	regs = append(regs, 1, 2, 3)
	fmt.Println(adapter(regs))
}

//数组转指定格式
func adapter(params []interface{}) string {
	var regs string
	argslen := len(params)
	//var isTypeEq bool = true
	for i, value := range params {
		//	k := reflect.TypeOf(params[i]).Kind()

		if i != argslen-1 {
			regs += strconv.Itoa(int(reflect.ValueOf(value).Int())) + ","
		} else {
			regs += strconv.Itoa(int(reflect.ValueOf(value).Int()))
		}
	}
	return regs
}

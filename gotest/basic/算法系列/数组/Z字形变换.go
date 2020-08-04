package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ZChange("LEETCODEISHIRING", 3))
}

/*

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
L   C   I   R
E T O E S I I G
E   D   H   N

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。


*/
//重点是找到一个周期
//每一行的元素数据的规律
func ZChange(data string, nums int) string {
	if nums == 1 {
		return data
	}
	var b = []rune(data)
	var res = make([]string, nums)
	var length = len(b)
	var period = nums*2 - 2
	for i := 0; i < length; i++ {
		var mod = i % period
		if mod < nums {
			res[mod] += string(b[i])
		} else {
			res[period-mod] += string(b[i])
		}
	}
	return strings.Join(res, "")
}

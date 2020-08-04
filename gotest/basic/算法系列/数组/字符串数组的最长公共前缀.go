package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaxLength([]string{"abcde", "abcdefggad", "abcqqq"}))
}

/*
编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
示例1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""

*/
func MaxLength(data []string) string {
	if len(data) <= 1 {
		return ""
	}
	prefix := data[0]
	for _, res := range data {
		for strings.Index(res, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

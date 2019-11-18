package main

import "fmt"

func main() {
	s := []interface{}{"hello", "world", "hello", "golang", "hello", "ruby", "php", "java"}

	fmt.Println(removeDuplicateElement(s))
}

//通过使用map对slice去重
func removeDuplicateElement(addrs []interface{}) []interface{} {
	result := make([]interface{}, 0, len(addrs))
	temp := map[interface{}]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

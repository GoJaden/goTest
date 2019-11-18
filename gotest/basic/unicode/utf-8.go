package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println(utf8.MaxRune)

	s := "fldjl房间打开了家乐福健康"
	fmt.Println(utf8.ValidString(s))
}

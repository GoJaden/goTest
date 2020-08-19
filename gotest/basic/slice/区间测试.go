package main

import "fmt"

func main() {
	a := "123456789"
	b := a[3:8]
	//切片的末尾位置是切片的容量
	fmt.Println(b)

}

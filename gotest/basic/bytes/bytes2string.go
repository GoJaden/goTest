package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func BytesToString(b []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	strHeader := reflect.StringHeader{bytesHeader.Data, bytesHeader.Len}
	return *(*string)(unsafe.Pointer(&strHeader))
}

func main() {
	bts := []byte("fjdkjlfskjdklffffffffjlfjdklfjlk就反了空当接龙开房记录弗兰克简单看了附近的考虑吉林大街李开复建档立卡房间里看到荆防颗粒几点几分考虑到健康了解来看")
	now := time.Now().UnixNano()
	for i := 0; i < 10000000; i++ {
		//_ = string(bts)
		BytesToString(bts)
	}

	fmt.Println(time.Now().UnixNano() - now)
}

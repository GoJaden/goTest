package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

//goroutine结束的时候是goroutine对应的函数结束时
func main() {
	marshalerType := reflect.TypeOf((*int)(nil)).Elem()
	fmt.Println(marshalerType)
	testType(132)
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("j: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func testType(i interface{}) {
	dd := i.(string)
	fmt.Println(dd)
	if s, ok := i.(string); !ok {
		fmt.Println("不是字符串类型")
	} else {
		fmt.Println(s)
	}
}

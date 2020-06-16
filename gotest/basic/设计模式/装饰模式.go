package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	for i := 0; i < 100; i++ {
		WrapperNumFunc(Number2, log.New(os.Stdout, "JADEN", 1))(i)
	}

}

//装饰模式,对一个函数进行加强
//常见的案例  gin 封装中间件，记录请求日志
type NumFunc func(a int) int

func WrapperNumFunc(fun NumFunc, logger *log.Logger) NumFunc {
	return func(a int) int {
		fn := func(b int) (result int) {
			defer func() {
				logger.Printf("num=%v,result=%v\n", b, result)
			}()
			fmt.Println("执行计算开始")
			return fun(b)
		}
		return fn(a)
	}
}

//计算一个数的平方
func Number2(a int) int {

	return a * a
}

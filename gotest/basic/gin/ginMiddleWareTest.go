package main

import (
	"common-lib/errcode"
	"common-lib/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

var defaultEngine *gin.Engine

func main() {

	defaultEngine = gin.New()

	addMiddleWare(loggerEngine())

	defaultEngine.GET("/test", func(c *gin.Context) {
		panic("213")
		fmt.Println("--------------------------------------+++++++++++++++++++++++++++")
		c.Writer.WriteString("hello")
	})
	defaultEngine.Run(":8082")
	fmt.Println(defaultEngine)
}

func loggerEngine() func(*gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			//拦截全局异常，并打印日志
			if err := recover(); err != nil {
				fmt.Printf("拦截到异常，开始打印输出")
				log := logger.DefaultSudaLogger
				log.Error(logger.LogInterface, c.Request.URL.Path, logger.LogResult,
					fmt.Sprintf("panic:%v", errors.Wrap(err, 2).ErrorStack()),
					logger.LogAlarmID, "1000000000000000ID",
				)
				c.JSON(200, errcode.SYSTEM_ERROR.ToResult())
			}

		}()
		//执行链的下一步任务
		c.Next()
	}
}

func addMiddleWare(middleWare func(*gin.Context)) {
	defaultEngine.Use(middleWare)
}

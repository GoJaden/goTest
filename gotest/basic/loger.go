package main

import (
	"fmt"
)

type Logger interface {
	//写入数据的方法，返回成功与否
	writeData(interface{}) error
}

type LoggerWriter struct {
	//利用切面存放所有的实现logger接口的实例
	writers []Logger
}

func (log *LoggerWriter) addLogger(logger Logger) {
	//这个地方必须重新赋值才行
	log.writers = append(log.writers, logger)
}

func (l *LoggerWriter) writeData(data interface{}) error {
	for _, v := range l.writers {
		v.writeData(data)
	}
	return nil
}

func main() {
	fmt.Println("---")
	fw := new(FileWriter)

	lw := new(LoggerWriter)
	lw.addLogger(fw)
	lw.writeData("开始记录日志")
}

type FileWriter struct {
}

func (f *FileWriter) writeData(data interface{}) error {
	fmt.Println("写入日志到文件中...", data)
	return nil
}

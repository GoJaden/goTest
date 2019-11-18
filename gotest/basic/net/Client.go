package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("客户端连接服务端出现意外:", err)
		return
	}
	code, err := conn.Write([]byte("你好呀，我是客户端，我有点话跟你说"))
	if err != nil {
		fmt.Println("意外:", err)
		return
	}
	fmt.Println("客户端发送数据，状态码:", code)
	var buffer []byte = make([]byte, 1024)
	code2, err := conn.Read(buffer)
	fmt.Println(code2, string(buffer))
}

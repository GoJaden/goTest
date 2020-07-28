package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	server()
}

func server() {

	lister, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("开启服务端失败")
		return
	}

	for {
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println("接收客户端请求失败")
			break
		}
		go DealConn(conn)

	}
}

func DealConn(conn net.Conn) {
	timer := time.NewTicker(time.Second)

	for {
		select {
		case <-timer.C:
			{
				content := make([]byte, 60)
				conn.Read(content)

				fmt.Printf("正在处理客户端%s的请求,处理内容是:%s", conn.RemoteAddr().String(), string(content))

				_, err := conn.Write([]byte("已经收到了客户端发送的处理请求,谢谢"))
				if err != nil {
					fmt.Println(err)
				}
			}
		}

	}

}

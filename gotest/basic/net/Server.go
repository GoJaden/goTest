package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务端即将启动...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("服务端启动错误，错误信息是:", err)
		return
	}

	fmt.Printf("客户端地址是:%s,客户端网络是：%s\n", listener.Addr().String(), listener.Addr().Network())
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受客户端请求的时候异常了:", err)
			return
		}

		go handlerInfo(conn)
	}
}

func handlerInfo(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("loacladdr:%s,remoteaddr:%s\n", conn.LocalAddr(), conn.RemoteAddr())
	var buffer []byte = make([]byte, 1024)
	code, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("读取客户端的数据出了问题:", err)
	}
	fmt.Println("code:", code)
	fmt.Println("客户端发送的数据是：", string(buffer))

	conn.Write([]byte("我是服务端，我已经接受到了数据，88"))

}

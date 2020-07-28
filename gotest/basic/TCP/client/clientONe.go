package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	clientSend()
}

func clientSend() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("建立连接失败")
		return
	}

	tiemr := time.NewTicker(time.Second * 3)

	for {
		select {
		case <-tiemr.C:
			{
				fmt.Println("定时触发")
				_, err := conn.Write([]byte("你好呀，我是客户端~~~~~~"))
				if err != nil {
					fmt.Println(err)
					return
				}
				bts := make([]byte, 60)

				_, err = conn.Read(bts)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("服务端给我响应的数据是:" + string(bts))
			}

		}

	}

}

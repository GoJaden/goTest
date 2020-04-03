package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	var count int

	service := ":12000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError1(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError1(err)

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}
		count++
		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		//conn.Close()                // we're finished with this client
		fmt.Println("连接数量:", count)
	}
}
func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

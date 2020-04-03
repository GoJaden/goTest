package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {

	//IP_MASK()

	//HostToIp()
	Dail()
	//LookHost()
}

//tcp通信
func Dail() {
	for {

		service := "127.0.0.1:12000"

		tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
		checkError(err)

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		conn.SetKeepAlive(true)
		fmt.Println(conn)
		checkError(err)

		//_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
		checkError(err)

		//result, err := readFully(conn)
		result, err := ioutil.ReadAll(conn)
		checkError(err)

		fmt.Println(string(result))
		time.Sleep(time.Second * 3)
	}

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//查询主机
func LookHost() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}

//解析域名对应的ip地址
func HostToIp() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())
	os.Exit(0)
}

//获取ip 、网关地址
func IP_MASK() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		" Default mask length is ", bits,
		"Leading ones count is ", ones,
		"Mask is (hex) ", mask.String(),
		" Network is ", network.String())
	os.Exit(0)
}

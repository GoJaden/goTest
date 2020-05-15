package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	cmd := exec.Command("ls")
	reader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("获取读取流失败")
		return
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println("run err")
		return
	}

	read := bufio.NewReader(reader)
	fmt.Println("开始读取")
	if data, _, err := read.ReadLine(); err != io.EOF {

		fmt.Println(string(data))
	}

	cmd.Wait()
	fmt.Println("run success")
}

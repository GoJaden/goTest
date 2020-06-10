package main

import (
	"fmt"
	"github.com/go-stack/stack"
	"io/ioutil"
	"os/exec"
)

func main() {
	t()
	c := stack.Caller(0)

	fmt.Println(c.String())

}
func t() {
	e := exec.Command("go", "run", "E:\\goProject\\src\\goProject\\gotest\\basic\\test\\for.go")

	out, err := e.StdoutPipe()
	if err != nil {
		fmt.Println("out:", err)
		return
	}

	eout, err := e.StderrPipe()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	err = e.Start()
	if err != nil {
		fmt.Println("start:", err)
		return
	}

	bts, err := ioutil.ReadAll(out)
	if err != nil {
		fmt.Println("read:", err)
		return
	}

	bts1, err := ioutil.ReadAll(eout)
	if err != nil {
		fmt.Println("read1:", err)
		return
	}
	fmt.Println("result:", string(bts), string(bts1))
	e.Wait()

}

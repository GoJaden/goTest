package main

import (
	"fmt"
	"os"
)

func main() {

	logFile, err := os.OpenFile("../../1.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(logFile)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go sendData(channel)

	time.Sleep(time.Second * 3)
	go getData(channel)
}

func sendData(channel1 chan string) {
	channel1 <- "a"
	channel1 <- "b"
	channel1 <- "c"
	channel1 <- "d"
	channel1 <- "e"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s", input)
	}
}

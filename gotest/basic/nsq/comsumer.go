package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

func sendMessage(publisher *nsq.Producer) {

	if err := publisher.Publish("jaden", []byte("111111111111111")); err != nil {
		panic(err)
	}

}
func sendMessage1(publisher *nsq.Producer) {

	if err := publisher.Publish("jaden", []byte("2222222222222222")); err != nil {
		panic(err)
	}

}

func main() {
	url := "192.168.1.204:4150"
	publisher, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		fmt.Print(err)
	}
	go sendMessage(publisher)
	go sendMessage1(publisher)
	time.Sleep(time.Second * 2)
	publisher.Stop()
}

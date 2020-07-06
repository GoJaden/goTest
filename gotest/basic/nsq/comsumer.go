package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

func sendMessage(publisher *nsq.Producer) {

	if err := publisher.Publish("jaden", []byte("12fdjkfjlkdfds")); err != nil {
		panic(err)
	}

}
func sendMessage1(publisher *nsq.Producer) {

	if err := publisher.Publish("jaden", []byte("12fdjk32132132fjlkdfds")); err != nil {
		panic(err)
	}

}

func main() {
	url := "127.0.0.1:4150"
	publisher, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		fmt.Print(err)
	}
	sendMessage(publisher)
	sendMessage1(publisher)

	publisher.Stop()
}

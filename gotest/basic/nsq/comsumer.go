package main

import (
	"common-lib/nsq"
	"time"
)

func sendMessage(publisher *nsq.Producer) {

	if err := publisher.Publish("test2", []byte("jaden到此一游")); err != nil {
		panic(err)
	}

}
func sendMessage1(publisher *nsq.Producer) {

	if err := publisher.Publish("test3", []byte("test3333333333333")); err != nil {
		panic(err)
	}

}

func main() {
	url := "192.168.1.204:4150"
	publisher := nsq.NewProducer(url)
	for i := 0; i < 100; i++ {
		sendMessage(publisher)
	}

	for i := 0; i < 100; i++ {
		sendMessage1(publisher)
	}
	time.Sleep(time.Second * 5)
	publisher.Stop()
}

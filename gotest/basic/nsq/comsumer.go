package main

import (
	"common-lib/nsq"
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

	sendMessage(publisher)
	sendMessage1(publisher)

	publisher.Stop()
}

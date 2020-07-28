package main

import (
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"sync"
)

func main() {
	testNSQ()
}

type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return errors.New("error-----------")
}

func testNSQ() {

	url := "192.168.1.204:4150"

	waiter := sync.WaitGroup{}
	waiter.Add(1)

	defer waiter.Done()
	config := nsq.NewConfig()
	config.MaxInFlight = 9

	consumer1, err := nsq.NewConsumer("jaden", "123", config)
	if nil != err {
		fmt.Println("err", err)
		return
	}

	err = consumer1.ConnectToNSQD(url)
	if nil != err {
		fmt.Println("err111", err)
		return
	}
	consumer1.AddHandler(&NSQHandler{})

	waiter.Wait()
}

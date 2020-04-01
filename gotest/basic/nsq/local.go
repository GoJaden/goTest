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


			consumer, err := nsq.NewConsumer("test3", "123", config)
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD(url)
			if nil != err {
				fmt.Println("err", err)
				return
			}
		consumer1, err := nsq.NewConsumer("test2", "123", config)
		if nil != err {
			fmt.Println("err", err)
			return
		}

		consumer1.AddHandler(&NSQHandler{})
		err = consumer1.ConnectToNSQD(url)
		if nil != err {
			fmt.Println("err", err)
			return
		}


	waiter.Wait()
}

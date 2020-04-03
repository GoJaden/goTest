package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	c := sync.NewCond(&m)
	n := 2
	running := make(chan bool, n)
	awake := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func() {
			m.Lock()
			fmt.Println("发送信号running")
			running <- true
			c.Wait()

			awake <- true
			m.Unlock()
		}()
	}
	//time.Sleep(time.Second*1)
	for i := 0; i < n; i++ {
		bool := <-running // Wait for everyone to run.
		fmt.Println(bool)
	}
	for n > 0 {
		select {
		case <-awake:
			fmt.Println("goroutine not asleep")
		default:
			fmt.Println("---goroutine not asleep")
		}
		m.Lock()
		c.Signal()
		m.Unlock()
		wake := <-awake // Will deadlock if no goroutine wakes up
		fmt.Println("wake:", wake)
		select {
		case <-awake:
			fmt.Println("too many goroutines awake")
		default:
		}
		n--
	}
	c.Signal()
}

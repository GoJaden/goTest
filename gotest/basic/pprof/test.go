package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("jha")
		f1()
		f2()
		writer.Write([]byte("hello "))
	})
	log.Println(http.ListenAndServe(":8080", nil))
}

func f1() {

	i := 0
	for i < 1000000 {
		i++
	}
	time.Sleep(time.Second)

}

func f2() {
	fmt.Println("f2")
	time.Sleep(time.Second * 3)
}

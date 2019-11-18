package main

import (
	"bufio"
	"bytes"
	"fmt"
	"time"
)

func main() {
	Hello()

	var a [5]int
	a[0] = 1
	fmt.Println(a[0])

	s := make([]int, 3)
	s[0] = 1
	s[2] = 2
	fmt.Print(s[0:])
	fmt.Println(len(s))
	ss := append(s, 1, 2, 3, 4, 5)
	fmt.Println(ss)

	d := make([]int, len(ss))
	copy(d, ss)
	fmt.Println(d)

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)
	delete(m, "a")
	fmt.Println(m)

	n := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println(n)
	fmt.Println(plus(22, 11))
	_, bb := plus2(22, 11)
	fmt.Println(bb)

	fun := add()
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())

	fmt.Println(expr(3))

	per := person{"jiangjiu", 23}
	per.age = 30
	fmt.Println(per)

	r := &Rect{4, 5}
	fmt.Println(MesureArea(r))
	fmt.Println(MesurePerim(r))

	fmt.Println(Exists(2))

	/*f("direct")

	go f("gorouting")

	go func(s string) {
		fmt.Println(s)
	}("going")
	fmt.Scanln()
	go func() {
		fmt.Println("123" +
			"")
	}()
	fmt.Println("done")*/

	/*messages := make(chan string,2)
	go func() {
		messages <- "hello chan"
		messages <-"abc"
		fmt.Println("----")
		messages <-"abc"
		messages <-"abc"
	}()

	fmt.Println(<- messages)
	fmt.Println(<- messages)
	fmt.Println(<- messages)
	fmt.Println(<- messages)*/

	/*done := make(chan bool ,1)
	go worker(done)

	fmt.Println(<-done)
	*/
	/*pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)


	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		c1<-"c1,hello"
	}()
	go func() {
		time.Sleep(time.Second*2)
		c2<-"c2,hello"
	}()
	for i := 0; i < 2; i++ {
		select {
			case msg :=<-c1:{
				fmt.Println(msg)
			}
			case msg :=<-c2:{
				fmt.Println(msg)
			}
		}
	}
	*/

	jobs := make(chan string, 2)
	jobs <- "msg"
	jobs <- "msg"
	close(jobs)

	for msg := range jobs {
		fmt.Println(msg)
	}
	/*
		//定时器
		timer1 :=time.NewTimer(time.Second*5)
		fmt.Printf("%V\n",timer1)
		<-timer1.C
		fmt.Print("多少时间后，我开始执行111")

		//打点器
		ticker1 := time.NewTicker(time.Second*1)

		go func() {
			for t := range ticker1.C{
				fmt.Println(t)
			}
		}()

		time.Sleep(time.Second*10)
		ticker1.Stop()
		fmt.Println("10s后结束打点任务")*/

	/*var ops uint64 =0

	for i := 0; i < 50; i++ {
		go func() {
			atomic.AddUint64(&ops,1)
			runtime.Gosched()
		}()
	}
	//time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println(opsFinal)




	var aa interface{} = "32"
		fmt.Printf("%T",aa)



	o := new(object)

	fmt.Printf("%T\n",o)

	aaa:= []int{1,23,3}

	result := Result{}
	result.Data=aaa
	result.Msg="成功请求"
	result.Status=1

	res,err := json.Marshal(result)
	if err==nil{
		fmt.Println(string(res))
	}

	json.Unmarshal(res,result)
	fmt.Println(result)

	cons := new (Consumer)
	prod := new (Producer)

	r1 := make(chan string,1)
	go prod.product("test",r1)
	//go cons.consum(r1) 这样子写，会将阻塞操作转换成非阻塞的操作，导致程序结束
	cons.consum(r1)*/

	bbb := new(byte)
	fmt.Println(bbb)
	/*
		l := list2.New()
		l.PushBack("fdf")
		l.PushBack(12)
		for i:=l.Front();i!=nil;i=i.Next(){
			fmt.Println(*i)
		}

		http.Handle("/",http.FileServer(http.Dir(".")))
		http.ListenAndServe(":8080",nil)
	*/

	//io
	/*data := []byte("中华人民共和国")
	rd := bytes.NewReader(data)
	rrr := bufio.NewReader(rd)
	var buf [128]byte
	nn, err := rrr.Read(buf[:])
	fmt.Println(string(buf[:nn]), nn,err)*/
	str := "hello"
	data := []byte(str)

	read := bytes.NewReader(data)
	rea := bufio.NewReader(read)
	var buffer []byte = make([]byte, 5)
	fmt.Println(len(buffer))
	va, err := rea.Read(buffer[:])
	if err == nil {
		fmt.Println("打印的值是:", va, string(buffer[0:]))
	}

}

type Consumer struct {
}
type Producer struct {
}

func (p Producer) product(info string, channel chan<- string) {
	for {
		channel <- info
		time.Sleep(time.Second)
	}
}

func (c Consumer) consum(channel <-chan string) {

	for {
		info := <-channel
		fmt.Println(info)
	}
}

type Result struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

type myInter interface {
	gothing() int
}

type object struct {
	name string
}

func (o object) gothing() int {

	return 1
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func worker(done chan bool) {
	fmt.Println("starting...")
	time.Sleep(time.Second * 20)
	fmt.Println("done")
	done <- true
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func expr(i int) int {
	if i == 0 {
		return 1
	}
	return i * expr(i-1)
}

func plus(a, b int) int {
	c := a + b
	return c
}

func plus2(a, b int) (int, int) {
	c := a + b
	return c, c
}

func add() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

type person struct {
	name string
	age  int
}

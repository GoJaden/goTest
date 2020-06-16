package main

import "fmt"

func main() {
	fmt.Println(NewProxy(&Person{}).Do("我需要租房子"))
}

//静态代理
type Task interface {
	Do(something string) string
}

type Person struct {
}

func (p *Person) Do(something string) string {
	fmt.Println("开始做事情" + something)
	return fmt.Sprintf("p=%+v,已经做完了", p)
}

type Proxy struct {
	Task
}

func NewProxy(t Task) *Proxy {
	return &Proxy{t}
}

func (p *Proxy) ProxyDo(something string) {
	p.Do(something)
}

//动态代理

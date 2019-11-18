package main

import "fmt"

type Context struct {
	Id   int
	Name string
}

type DealFunc func(c *Context)

func (dd DealFunc) resolve() {
	fmt.Println("------")
}
func RealDeal(d DealFunc) {
	d.resolve()
}

func FuncOne(ctx *Context) {
	fmt.Println(ctx)
}
func main() {
	ctx := new(Context)
	ctx.Id = 100
	var a DealFunc = FuncOne
	a(ctx)
	RealDeal(a)

}

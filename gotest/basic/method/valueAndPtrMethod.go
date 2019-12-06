package main // import "text"

import "fmt"

type Animal interface {
	Bark()
}

type Dog struct {
}

func (d Dog) Bark() {
	fmt.Println("dog")
}

type Cat struct {
}

func (c *Cat) Bark() {
	fmt.Println("cat")
}

func Bark(a Animal) {
	a.Bark()
}

func getDog() Dog {
	return Dog{}
}

func getCat() Cat {
	return Cat{}
}

//如果是值方法，指针或者值都可以调用(go里面有个机制，将地址转化为具体的值)
//	如果是指针方法，则只能指针调用
/*对于类型T，它的方法集合是所有接收者为T的方法。
对于类型*T，它的方法集合是所有接收者为*T和T的方法。*/
func main() {
	dp := &Dog{}
	d := Dog{}
	dp.Bark() // (1)
	d.Bark()  // (2)
	Bark(dp)  // (3)
	Bark(d)   // (4)

	cp := &Cat{}
	c := Cat{}
	cp.Bark() // (5)
	c.Bark()  // (6)
	Bark(cp)  // (7)
	Bark(&c)  // (8)

	getDog().Bark() // (9)
	//getCat().Bark() // (10)
}

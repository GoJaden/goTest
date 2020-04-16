package main

import (
	"fmt"
	"time"
)

func main() {
	p := &Person{
		Animal: Animal{
			Desc: "这是动物的描述信息",
			Time: time.Now(),
		},
		Name: "jaden",
		Age:  23,
	}
	p.Eat("牛排")
	a := Animal{
		Desc: "一只宠物猫",
		Time: time.Now(),
	}
	p.Play(a)
	p.Speak()

}

type doing interface {
	//吃东西
	Eat(string)
	//跟别的动物玩
	Play(animal Animal)
	//说话
	Speak() string
}

func (a Animal) Eat(foodName string) {
	fmt.Printf("动物%s正在吃%s\n", a, foodName)
}
func (a Animal) Play(animal Animal) {
	fmt.Printf("动物%s正在跟动物%s玩\n", a, animal)
}
func (a Animal) Speak() string {
	return fmt.Sprintf("动物%s在打招呼说,hello", a)
}

func (a Person) Eat(foodName string) {
	fmt.Printf("动物%s正在吃%s\n", a, foodName)
}
func (a Person) Play(animal Animal) {
	fmt.Printf("动物%s正在跟动物%s玩\n", a, animal)
}
func (a Person) Speak() string {
	return fmt.Sprintf("动物%s在打招呼说,hello", a)
}

type Animal struct {
	Desc string
	Time time.Time
}

type Person struct {
	Animal
	Name string
	Age  int
}

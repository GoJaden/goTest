package main

import "fmt"

func main() {
	m := NewMachine()
	m.On()
	m.Off()
	m.On()
	m.Off()
	m.Off()
}

//状态模式  减少if else 判断
//通过不同的状态，实现不同的功能

type State interface {
	On(m *machine)
	Off(m *machine)
}

type machine struct {
	current State
}

func NewMachine() *machine {
	return &machine{
		current: NewOffState(),
	}
}
func (m *machine) On() {
	m.current.On(m)
}
func (m *machine) Off() {
	m.current.Off(m)
}

type OnState struct {
}

func NewOnState() *OnState {
	return new(OnState)
}

func (on *OnState) On(m *machine) {
	fmt.Println("机器已经开启了...")
}
func (on *OnState) Off(m *machine) {
	fmt.Println("关闭了机器")
	m.current = NewOffState()
}

type OffState struct {
}

func NewOffState() *OffState {
	return new(OffState)
}
func (off *OffState) On(m *machine) {
	fmt.Println("开启机器")
	m.current = NewOnState()
}
func (off *OffState) Off(m *machine) {
	fmt.Println("机器已经关闭...")
}

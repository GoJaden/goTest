package main

import "errors"

type Geometry interface {
	Area() int
	Perim() int
}

func Area() int {
	return 1
}

type Rect struct {
	Width  int
	Height int
}

func (r *Rect) Area() int {
	return r.Height * r.Width
}

func (r *Rect) Perim() int {
	return 2*r.Width + 2*r.Height
}

func MesureArea(g Geometry) int {
	return g.Area()
}

func MesurePerim(g Geometry) int {
	return g.Perim()
}

func Exists(a int) (int, error) {
	if a == 1 {
		return 1, errors.New("存在对应的值...")
	}
	return 0, errors.New("不存在对应的值...")
}

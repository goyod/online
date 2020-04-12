package interfaces

import "fmt"

type Areaer interface {
	Area() float64
}

func Area(a Areaer) {
	fmt.Println(a.Area())
}

type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(base, height float64) Triangle {
	return Triangle{
		base:   base,
		height: height,
	}
}

func (r Triangle) Area() float64 {
	return r.base * r.height * 0.5
}

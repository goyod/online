package structure

import "fmt"

type rectangle struct {
	w, l float64
}

func Construct() {
	rec := rectangle{w: 10, l: 12}

	area := rec.w * rec.l

	fmt.Printf("area of rectangle %fx%f = %f\n", rec.l, rec.w, area)
}

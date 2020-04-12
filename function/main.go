package main

import (
	"fmt"
	"function/interfaces"
	"function/method"
)

func main() {
	rec := method.NewRectangle(4, 5)

	fmt.Println(rec.Area())

	rec.SetWidth(12)
	fmt.Println("it's change:", rec.Area())

	rec.SetLength(50)
	fmt.Println("it's not change:", rec.Area())

	tri := interfaces.NewTriangle(9, 10)
	interfaces.Area(tri)
	interfaces.Area(rec)
}

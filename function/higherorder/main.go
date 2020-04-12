package main

import "fmt"

func main() {
	higherOrderParams(area)
	fmt.Println(higherOrderReturns()(3, 4))
}

func area(dx, dy float64) float64 {
	return dx * dy
}

func higherOrderParams(firstClass func(float64, float64) float64) {
	fmt.Printf("area of 64x48 = %.0f", firstClass(64, 48))
}

func higherOrderReturns() (firstClass func(float64, float64) float64) {
	return area
}

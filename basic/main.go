package main

import (
	"fmt"
)

func main() {
	variables()
}

func variables() {
	{
		var s string
		s = "สวัสดี"
		fmt.Println(s)
	}

	{
		var s string = "Hello World"
		fmt.Println(s)
	}

	{
		var i int
		i = 14
		fmt.Printf("%d\n", i)
	}

	{
		var i int = 9
		fmt.Printf("%v\n", i)
	}

	{
		i := 9
		fmt.Println(i)
	}
}

func add(a int, b int) int {
	return a + b
}

func ifelse() {
	a, b := 1, 2

	if a != b {
		println("a not equal to b")
	} else if a < b {
		println("a less than b")
	} else {
		println("ok")
	}
}

func ifWithStatement() {
	if a := false; a != true {
		fmt.Println("a is not true")
	}
}

func forloop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for {
		return
	}
}

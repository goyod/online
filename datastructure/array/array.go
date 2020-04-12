package array

import "fmt"

func ListIndex() {
	var array [5]string

	for i := range array {
		fmt.Println(i)
	}
}

func ListIndexValue() {
	var array [5]string

	for i, v := range array {
		fmt.Printf("%d %q\n", i, v)
	}
}

func ListValue() {
	var array [5]string

	for _, v := range array {
		fmt.Printf("%q\n", v)
	}
}

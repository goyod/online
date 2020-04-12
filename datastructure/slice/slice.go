package slice

import "fmt"

func List() {
	var slice []int

	primes := [...]int{0, 1, 1, 2, 3, 5, 8, 13}
	slice = primes[3:7]

	for i := range slice {
		fmt.Println(slice[i])
	}
}

func MakeSlice() {
	slice := make([]string, 3, 3)
	slice[0] = "apple"
	slice[1] = "banana"
	slice[2] = "coconut"

	for _, v := range slice {
		fmt.Println(v)
	}
}

func AppendSlice() {
	slice := []string{}
	slice = append(slice, "apple", "banana", "coconut")
	slice = append(slice, "durian")

	for _, v := range slice {
		fmt.Println(v)
	}
}

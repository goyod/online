package maps

import "fmt"

func Make() {
	m := make(map[string]string)

	if m == nil {
		fmt.Println("it's nil")
	}

	m["a"] = "apple"
	m["b"] = "banana"
	m["c"] = "coconut"
	m["d"] = "durian"
	m["e"] = "elderberry"
	m["f"] = "fig"
	m["g"] = "guava"

	for k, v := range m {
		fmt.Printf("%q: %q\n", k, v)
	}
}

func Initial() {
	m := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "coconut",
		"d": "durian",
		"e": "elderberry",
		"f": "fig",
		"g": "guava",
	}

	delete(m, "d")

	for k, v := range m {
		fmt.Println(k, v)
	}
}

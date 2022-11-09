package main

import "fmt"

func main() {
	var i interface{}
	i = int(42)

	a, ok := i.(int)
	// a == 42 and ok == true

	fmt.Printf("a: %2d - b: %t\n", a, ok)

	b, ok := i.(string)
	// b == "" (default value) and ok == false

	fmt.Printf("a: %2s - b: %t\n", b, ok)
}

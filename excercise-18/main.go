package main

import "fmt"

var result int

func main() {
	fmt.Println("Init")

	result = midOperations(add)(2, 3)

	fmt.Println(result)

	result = midOperations(sub)(10, 6)

	fmt.Println(result)

	result = midOperations(prod)(2, 4)

	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func prod(a, b int) int {
	return a * b
}

func midOperations(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Operation Init")

		return f(a, b)
	}
}

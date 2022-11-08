package main

import "fmt"

var Calculus func(int, int) int = func(num1 int, num2 int) int {
	return num1 * num2
}

func main() {
	fmt.Println(Calculus(5, 7))

	// Rest
	Calculus = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Rest: %d\n", Calculus(6, 4))

	Operations()

	fmt.Println()

	twoTable := 2
	myTable := Table(twoTable)

	for i := 1; i < 11; i++ {
		fmt.Println(myTable())
	}
}

func Operations() {
	result := func() int {
		var a int = 23
		var b int = 14

		return a + b
	}

	fmt.Println(result())
}

func Table(value int) func() int {
	number := value
	sequence := 0

	return func() int {
		sequence++

		return number * sequence
	}
}

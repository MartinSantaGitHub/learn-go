package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var number1 int
	var number2 int
	var result int
	var legend string

	fmt.Println("Enter number 1: ")
	fmt.Scanln(&number1)

	fmt.Println("Enter number 2: ")
	fmt.Scanln(&number2)

	fmt.Println("Description: ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		legend = scanner.Text()
	}

	result = number1 + number2

	fmt.Println(legend, result)
}

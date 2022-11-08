package main

import "fmt"

var status bool = true

func main() {
	if status = false; status {
		fmt.Println(status)
	} else {
		fmt.Println("NOT")
	}

	switch number := 1; number {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Nothing")
	}
}

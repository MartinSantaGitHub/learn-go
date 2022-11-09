package main

import "fmt"

func main() {
	b := 6

	// *int is used to declare variable
	// b_pointer to be a pointer to an int
	var b_pointer *int

	// b_pointer is assigned the value that is the
	// address of where variable b is stored
	b_pointer = &b

	// Shorthand for the above two lines is:
	// b_pointer := &b

	fmt.Printf("Address of b_pointer: %p\n", b_pointer)

	// We can use *b_pointer to get the value that is stored
	// at address b_pointer, known as dereferencing the pointer
	fmt.Printf("Value stored at b_pointer: %d\n", *b_pointer)

	// Same as:
	//fmt.Printf("Value stored at b_pointer: %d\n", *&b)
}

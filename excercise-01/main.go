package main

import (
	"fmt"
)

type User struct {
	Age int
}

func main() {
	var u *User

	defer func() {
		fmt.Println("Crash!")
	}()

	u.Age = 10

	var c []int

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	c = append(c, a...)
	c = append(c, b...)

	fmt.Print(c)

	//fmt.Println(u)
	fmt.Println("Hello World!")

	// data := make(map[string]string)

	// data["123"] = "Algo"

	// fmt.Println(data["123"])
}

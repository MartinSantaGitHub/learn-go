package main

import "fmt"

type User struct {
	Age int
}

func main() {
	var u *User

	defer func() {
		fmt.Println("Crash!")
	}()

	u.Age = 10

	fmt.Println(u)
	fmt.Println("Hello World!")
}

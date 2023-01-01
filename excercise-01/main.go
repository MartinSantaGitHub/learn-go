package main

import (
	"fmt"
)

type User struct {
	Age int
}

type Usuario struct {
	Id        string
	Name      string
	LastName  string
	Email     string
	Password  string
	Avatar    string
	Banner    string
	Biography string
	Location  string
	WebSite   string
}

func main() {
	// var u *User

	// defer func() {
	// 	fmt.Println("Crash!")
	// }()

	//u.Age = 10

	var c []int

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	c = append(c, a...)
	c = append(c, b...)

	fmt.Print(c)

	//fmt.Println(u)
	fmt.Println("Hello World!")

	// data := make(map[string]*Usuario)

	// data["123"] = &Usuario{Name: "Pepe"}

	// retrieve := data["321"]

	// fmt.Println(retrieve == nil)

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	results := list[5:7]

	fmt.Println(results)
}

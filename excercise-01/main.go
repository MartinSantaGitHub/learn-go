package main

import (
	"fmt"
)

type User struct {
	Age int
}

type Usuario struct {
	Id        string `json:"id"`
	Name      string `json:"name,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Banner    string `json:"banner,omitempty"`
	Biography string `json:"biography,omitempty"`
	Location  string `json:"location,omitempty"`
	WebSite   string `json:"webSite,omitempty"`
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

	data := make(map[string]*Usuario)

	data["123"] = &Usuario{Name: "Pepe"}

	retrieve := data["321"]

	fmt.Println(retrieve == nil)
}

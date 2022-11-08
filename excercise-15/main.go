package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go mySlowName("Martin Santamaria")

	fmt.Println("I'm Here")

	var x string

	fmt.Scanln(&x)
}

func mySlowName(name string) {
	chars := strings.Split(name, "")

	for _, char := range chars {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(char)
	}
}

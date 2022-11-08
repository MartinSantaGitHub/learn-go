package main

import (
	"fmt"
	"strconv"
)

var Number int
var text string
var status bool

func main() {
	number4, number5, text := 5, 6.7, "lala"

	var coin float32 = 0
	var coin2 float64 = 5.12345

	coin = float32(coin2)

	text = fmt.Sprintf("%.2f", coin)
	text = strconv.Itoa(int(coin))

	fmt.Println(Number)
	fmt.Println(number4)
	fmt.Println(number5)
	fmt.Println(text)

	ShowStatus()
}

func ShowStatus() {
	fmt.Println("this is", status)
}

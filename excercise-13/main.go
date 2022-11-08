package main

import (
	"fmt"
)

func main() {
	pow2(2)
}

func pow2(number int) {
	if number > 100000000 {
		return
	}

	fmt.Println(number)

	pow2(number * 2)
}

package main

import (
	"fmt"
	"time"
)

func testFor(n int) {
	for i := 0; i < n; i++ {
		//fmt.Println(i)
	}
}

func main() {
	init := time.Now()

	testFor(10000000)

	end := time.Now()
	delta := end.Sub(init).Milliseconds()

	fmt.Println(delta)
}

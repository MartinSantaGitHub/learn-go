package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n > 1 {
		return fibonacci(n-2) + fibonacci(n-1)
	}

	return n
}

func main() {
	init := time.Now()

	fibonacci(42)

	end := time.Now()
	delta := end.Sub(init).Milliseconds()

	fmt.Println(delta)
}

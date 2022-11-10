package main

import (
	"fmt"
	"time"
)

func loop(channel chan time.Duration) {
	init := time.Now()

	for i := 0; i < 10000000000; i++ {

	}

	end := time.Now()

	channel <- end.Sub(init)
}

func main() {
	channel1 := make(chan time.Duration)

	go loop(channel1)

	fmt.Println("I came to here")

	// Here is going to do an await
	duration := <-channel1

	fmt.Println(duration)
}

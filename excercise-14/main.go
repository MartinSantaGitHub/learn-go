package main

import (
	"fmt"
	"log"
	"os"
)

func something() {
	fmt.Println("Hey Jude!")
}

func panicExample(a int) {
	defer func() {
		reco := recover()

		if reco != nil {
			log.Fatalf("Panic generated error\n%v", reco)
		}
	}()

	if a == 1 {
		panic("Don't let me down")
	}
}

func main() {
	file := "test.txt"
	_, err := os.Open(file)

	defer something()

	if err != nil {
		fmt.Println("Error opening the file")
		//os.Exit(1)
	}

	panicExample(1)
}

package main

import "fmt"

func main() {
	dict := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	arr := [4]string{"Barcelona", "Real Madrid", "Chivas", "Boca Juniors"}

	for idx, value := range arr {
		fmt.Printf("index: %d - value: %s\n", idx, value)
	}

	for key, value := range dict {
		fmt.Printf("key: %15s - value: %d\n", key, value)
	}
}

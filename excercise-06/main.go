package main

import "fmt"

func main() {
	fmt.Println(one(5))

	_, status := two(1)

	//fmt.Println(number)
	fmt.Println(status)

	fmt.Println(Calculo(1, 3, 5))
}

func one(number int) (z int) {
	z = number * 3

	return
}

func two(number int) (int, bool) {
	if number == 1 {
		return 5, false
	}

	return 10, true
}

func Calculo(number ...int) int {
	total := 0

	for idx, num := range number {
		fmt.Printf("\n index %d", idx)

		total += num
	}

	fmt.Println()

	return total
}

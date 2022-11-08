package main

import "fmt"

var table [10]int
var matrix [5][7]int

func main() {
	// table := [10]int{1, 2, 3, 4}

	// table[0] = 1
	// table[5] = 4

	// for i := 0; i < len(table); i++ {
	// 	fmt.Println(table[i])
	// }

	matrix := []int{2, 3, 5}

	//matrix[3][5] = 1

	fmt.Println(matrix)

	slice()
	slice2()
	slice3()
}

func slice() {
	elements := [5]int{1, 2, 3, 4, 5}
	portion := elements[1:3]

	fmt.Println(portion)
}

func slice2() {
	elements := make([]int, 5, 20)

	fmt.Printf("Length %d, Capacity %d\n", len(elements), cap(elements))

}

func slice3() {
	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Length %d, Capacity %d\n", len(nums), cap(nums))
}

package main

import "fmt"

func main() {

	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// number := 0

	// for {
	// 	fmt.Println("Continue")
	// 	fmt.Println("Secret Number")
	// 	fmt.Scanln(&number)

	// 	if number == 100 {
	// 		break
	// 	}
	// }

	// var i = 0

	// for i < 10 {
	// 	fmt.Printf("\n Value of i: %d", i)

	// 	if i == 5 {
	// 		fmt.Printf("\t\tmultiply by 2 \n")
	// 		i *= 2
	// 		continue
	// 	}

	// 	fmt.Printf("\t\tPass by here \n")
	// 	i++
	// }

	var i int = 0

SECTION:
	for i < 10 {
		if i == 4 {
			i += 2
			fmt.Println("Go to SECTION adding 2 to i.")
			goto SECTION
		}
		fmt.Printf("Value of i: %d\n", i)
		i++
	}
}

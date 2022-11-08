package main

import "fmt"

func main() {
	countries := make(map[string]string, 5)

	fmt.Println(countries)

	countries["Mexico"] = "D.F."
	countries["Argentina"] = "Buenos Aires"

	fmt.Println(countries)

	championship := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	championship["River Plate"] = 25

	fmt.Println(championship)

	delete(championship, "Boca Juniors")

	fmt.Println(championship)

	for team, score := range championship {
		fmt.Printf("The team %s, has a score of %d\n", team, score)
	}

	score, exists := championship["Mineiro"]

	fmt.Printf("The score captured is %d, and does the team exist? %t\n", score, exists)
}

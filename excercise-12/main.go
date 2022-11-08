package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func readFile() {
	file, err := ioutil.ReadFile("./file.txt")

	if err != nil {
		fmt.Println("Something went wrong")

		return
	}

	fmt.Println(string(file))
}

func readFile2(filepath string, verbose bool) int {
	file, err := os.Open(filepath)
	count := 0

	if err != nil {
		fmt.Println("Something went wrong")

		return 0
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		registry := scanner.Text()

		if verbose {
			fmt.Printf("Line > %s\n", registry)
		}

		count++
	}

	file.Close()

	return count
}

func saveFile() {
	file, err := os.Create("./newFile.txt")

	if err != nil {
		fmt.Println("Something went wrong")

		return
	}

	fmt.Fprint(file, "This is a new file\nThis is another one")

	file.Close()
}

func saveFile2(filepath string) {
	count := readFile2(filepath, false)

	count++

	if !Append(filepath, fmt.Sprintf("\nThis is another line %d", count)) {
		fmt.Println("Something went wrong")
	}
}

func Append(filepath string, text string) bool {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Something went wrong")

		return false
	}

	_, err = file.WriteString(text)

	if err != nil {
		fmt.Println("Something went wrong")

		return false
	}

	return true
}

func main() {
	//readFile()
	readFile2("./file.txt", true)
	//saveFile()
	saveFile2("./newFile.txt")
}

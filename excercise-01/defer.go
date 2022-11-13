package main

func execute(m string) {
	println(m)
}

func getValue(v int) int {
	defer execute("Hello World!")

	println("First line")

	println("Second line")

	println("Third line")

	return v
}

func main() {
	v := getValue(5)

	println(v)
}

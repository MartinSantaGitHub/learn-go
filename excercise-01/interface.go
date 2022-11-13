package main

import (
	"fmt"
	"io"
)

type Human struct {
	Name string
}

func (h Human) Close() error {
	h.Name = "Pepe"

	fmt.Println(h.Name)

	return nil
}

func main() {
	var h Human
	var v io.Closer

	v = h

	h.Close()
	v.Close()

	h = v.(Human)

	fmt.Println()

	h.Close()
	v.Close()
}

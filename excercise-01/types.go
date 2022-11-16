package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10 - 1

	fmt.Printf("%T\n", a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a).Kind())
}

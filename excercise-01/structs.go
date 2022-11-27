package main

import "fmt"

type animal interface {
	breathe()
	eat()
	isBreathing() bool
	isEating() bool
	description(add string) string
}

type man struct {
	breathing bool
	eating    bool
}

type dog struct {
	breathing bool
	eating    bool
}

func (m man) breathe() {
	m.breathing = true
}

func (m *man) eat() {
	m.eating = true
}

func (m man) isBreathing() bool {
	return m.breathing
}

func (m *man) isEating() bool {
	return m.eating
}

func (m man) description(add string) string {
	return "I'm a man - " + add
}

func (d dog) breathe() {
	d.breathing = true
}

func (d *dog) eat() {
	d.eating = true
}

func (d dog) isBreathing() bool {
	return d.breathing
}

func (d *dog) isEating() bool {
	return d.eating
}

func (d dog) description(add string) string {
	return "I'm a dog - " + add
}

func action(a animal, varType string) {
	a.description(varType)

	a.breathe()
	a.eat()

	fmt.Println(a.isBreathing())
	fmt.Println(a.isEating())
}

func main() {
	mref := new(man)

	action(mref, "ref")

	var d dog

	// The same in all cases
	(&d).breathe()
	(*&d).eat()

	fmt.Println(d.isBreathing())
	fmt.Println(d.isEating())
}

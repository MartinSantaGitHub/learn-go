package main

import "fmt"

type livingBeing interface {
	alive() bool
}

type human interface {
	breathe()
	think()
	eat()
	sex() string
	alive() bool
}

type animal interface {
	breathe()
	eat()
	carnivorous() bool
	alive() bool
}

type vegetable interface {
	classification() string
}

type man struct {
	age         int
	height      float32
	weight      float32
	isBreathing bool
	isThinking  bool
	isEating    bool
	isMan       bool
	isAlive     bool
}

type woman struct {
	man
}

type dog struct {
	isBreathing   bool
	isEating      bool
	isCarnivorous bool
	isAlive       bool
}

func (m *man) breathe() {
	m.isBreathing = true
}

func (m *man) eat() {
	m.isEating = true
}

func (m *man) think() {
	m.isThinking = true
}

func (m *man) sex() string {
	if m.isMan {
		return "Man"
	}

	return "Woman"
}

func (m *man) alive() bool {
	return m.isAlive
}

func (this *dog) breathe() {
	this.isBreathing = true
}

func (this *dog) eat() {
	this.isEating = true
}

func (d *dog) carnivorous() bool {
	return d.isCarnivorous
}

func (d *dog) alive() bool {
	return d.isAlive
}

func humanBreathing(h human) {
	h.breathe()

	fmt.Printf("I'm a %s and I'm breathing \n", h.sex())
}

func animalBreathing(a animal) {
	a.breathe()

	fmt.Printf("I'm an animal and I'm breathing \n")
}

func animalCarnivorous(a animal) int {
	if a.carnivorous() {
		return 1
	}

	return 0
}

func isAlive(l livingBeing) bool {
	return l.alive()
}

func main() {
	// pedro := new(man)

	// pedro.isMan = true

	// maria := new(woman)

	// humanBreathing(pedro)
	// humanBreathing(maria)

	totalCarnivorous := 0
	dogo := new(dog)

	dogo.isCarnivorous = true
	dogo.isAlive = true

	animalBreathing(dogo)

	totalCarnivorous += animalCarnivorous(dogo)
	totalCarnivorous += animalCarnivorous(dogo)

	fmt.Printf("Carnivorous Total %d\nAnimal Alive %t\n", totalCarnivorous, isAlive(dogo))
}

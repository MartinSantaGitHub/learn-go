package main

import "fmt"

type BankAccount struct {
	balance int
}

func accountAddressByVal(account BankAccount) {
	fmt.Println()
	fmt.Println("Inside accountAddressByVal")
	fmt.Println()
	fmt.Printf("Address of account:  %p\n", &account)
}

func accountAddressByRef(account *BankAccount) {
	fmt.Println()
	fmt.Println("Inside accountAddressByRef")
	fmt.Println()
	fmt.Printf("Address of account:  %p\n", account)
}

func main() {
	account := new(BankAccount)

	account.balance = 100

	fmt.Printf("Address of account:  %p\n", account)

	fmt.Println()

	// Each one of these lines creates a copy of account
	a := *account
	b := *account
	c := *account

	// Each one of them has different memory address
	fmt.Printf("Address of a:  %p\n", &a)
	fmt.Printf("Address of b:  %p\n", &b)
	fmt.Printf("Address of c:  %p\n", &c)

	// We modify the balance of "c" by its reference
	(&c).balance -= 20

	fmt.Println()

	fmt.Printf("Balance of c (ByRef):  %d\n", (&c).balance)
	fmt.Printf("Balance of c (ByVal):  %d\n", c.balance)

	fmt.Println()

	// This creates a copy of c
	d := *(&c)

	// This does the same, creates another copy of c
	f := c

	fmt.Printf("Address of d:  %p\n", &d)
	fmt.Printf("Address of f:  %p\n", &f)

	// This modifies the balance of "d" and "f" by their values (not a reference to them)
	d.balance -= 20
	f.balance -= 20

	// This obtains the reference of "c". "g" is a reference of "c"
	g := &c

	fmt.Println()

	fmt.Printf("Address of g (same as c):  %p\n", g)

	fmt.Println()

	// This modifies the balance of "c" by the reference of "g"
	g.balance -= 20

	fmt.Printf("Balance of c (ByRef):  %d\n", (&c).balance)
	fmt.Printf("Balance of c (ByVal):  %d\n", c.balance)

	// This creates a copy of "c". Remember, "g" is a reference of "c"
	h := *g

	fmt.Println()

	fmt.Printf("Address of h:  %p\n", &h)

	// Inside this method, a copy of "h" is created
	accountAddressByVal(h)

	// This method receives the reference of "h"
	accountAddressByRef(&h)
}

package main

type BankAccount struct {
	balance int
}

func accountAddressByVal(account BankAccount) {
	println()
	println("Inside accountAddressByVal")
	println()
	println("Address of account: ", &account)
}

func accountAddressByRef(account *BankAccount) {
	println()
	println("Inside accountAddressByRef")
	println()
	println("Address of account: ", account)
}

func main() {
	account := new(BankAccount)

	account.balance = 100

	println("Address of account: ", account)

	println()

	// Each one of these lines creates a copy of account
	a := *account
	b := *account
	c := *account

	// Each one of them has different memory address
	println("Address of a: ", &a)
	println("Address of b: ", &b)
	println("Address of c: ", &c)

	// We modify the balance of "c" by its reference
	(&c).balance -= 20

	println()

	println("Balance of c (ByRef): ", (&c).balance)
	println("Balance of c (ByVal): ", c.balance)

	println()

	// This creates a copy of c
	d := *(&c)

	// This does the same, creates another copy of c
	f := c

	println("Address of d: ", &d)
	println("Address of f: ", &f)

	// This modifies the balance of "d" and "f" by their values (not a reference to them)
	d.balance -= 20
	f.balance -= 20

	// This obtains the reference of "c". "g" is a reference of "c"
	g := &c

	println()

	println("Address of g (same as c): ", g)

	println()

	// This modifies the balance of "c" by the reference of "g"
	g.balance -= 20

	println("Balance of c (ByRef): ", (&c).balance)
	println("Balance of c (ByVal): ", c.balance)

	// This creates a copy of "c". Remember, "g" is a reference of "c"
	h := *g

	println()

	println("Address of h: ", &h)

	// Inside this method, a copy of "h" is created
	accountAddressByVal(h)

	// This method receives the reference of "h"
	accountAddressByRef(&h)
}

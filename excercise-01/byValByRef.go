package main

import (
	"fmt"
)

type BankAccount struct {
	balance int
}

/* This method receives a pointer to BankAccount */
func Withdraw(account *BankAccount, amount int) error {
	if account == nil {
		return fmt.Errorf("account argument is a nil pointer")
	}

	fmt.Println("\nWithdrawing ", amount)

	if account.balance < amount {
		return fmt.Errorf("Not enough money in the account")
	}

	account.balance = account.balance - amount

	return nil
}

/* This method receives a copy of BankAccount */
func BalanceAfterWithdraw(account BankAccount, amount int) int {
	fmt.Println("\nChecking balance after withdraw ", amount)

	account.balance = account.balance - amount

	return account.balance
}

func firstCase() {
	/* Here, we are creating an account value type */
	account := BankAccount{
		balance: 100,
	}

	fmt.Println("account =", account)

	/*
		Because the Withdraw method receives a pointer to a BankAccount, we need to pass
		the address of the memory where the BankAccount is allocated.
		Then, inside the method, the pointer is dereferenced and we can operate
		with the struct and its properties
	*/

	err := Withdraw(&account, 20)

	if err != nil {
		fmt.Println("error =", err)
	}

	fmt.Println("account =", account)

	/*
		We already have a value type of BankAccount. When we send this value in the method below
		a copy is created inside that method
	*/
	balance := BalanceAfterWithdraw(account, 20)

	fmt.Println("predicted balance =", balance)
	fmt.Println("account =", account)
}

func secondCase() {
	//var account *BankAccount

	/*
		Here, we are creating a BankAccount and the result
		is a reference to that BankAccount
	*/
	account := new(BankAccount)

	account.balance = 100

	//account = nil

	fmt.Println("account =", account)

	/*
		We already received a pointer when we explicty created the BankAccount
		in the variable account. So, we only pass that variable (that is a pointer)
		to the method
	*/
	err := Withdraw(account, 20)

	if err != nil {
		fmt.Println("error =", err)
	}

	fmt.Println("account =", account)

	/*
		This method receives a BankAccount value type. So, first, we dereference
		the account variable to obtain its content. Then, a copy of that content is
		received in the method BalanceAfterWithdraw
	*/
	balance := BalanceAfterWithdraw(*account, 20)

	fmt.Println("predicted balance =", balance)
	fmt.Println("account =", account)
}

func main() {
	fmt.Printf("First Case\n\n")
	firstCase()

	fmt.Printf("\nSecond Case\n\n")
	secondCase()
}

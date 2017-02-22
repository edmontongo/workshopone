package main

import "fmt"

// START TYPE OMIT

type AccountType uint

const (
	Checking AccountType = iota
	Savings
)

type Account struct {
	Name    string
	Balance float64
	AccountType
}

// END TYPE OMIT

// START PRINT OMIT

func printType(t AccountType) string {
	if t == Checking {
		return "checking"
	} else {
		return "savings"
	}
}

func Print(a Account) {
	s := fmt.Sprintf("%s has a %s account with $%.2f",
		a.Name, printType(a.AccountType), a.Balance)
	fmt.Println(s)
}

// END PRINT OMIT

// START WITHDRAW OMIT

func Withdraw(a *Account, amount float64) {
	a.Balance -= amount
}

func Deposit(a *Account, amount float64) {
	a.Balance += amount
}

// END WITHDRAW OMIT

func main() {
	// START PRINT MAIN OMIT
	a := Account{
		Name:        "Strans",
		Balance:     22.50,
		AccountType: Checking,
	}
	Print(a)
	// END PRINT MAIN OMIT

	// START WITHDRAW MAIN OMIT
	Withdraw(&a, 10.50)
	Print(a)

	Deposit(&a, 500)
	Print(a)
	// END WITHDRAW MAIN OMIT
}

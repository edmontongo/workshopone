package main

import "fmt"

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

// START PRINT OMIT

func (t AccountType) print() string {
	if t == Checking {
		return "checking"
	} else {
		return "savings"
	}
}

func (a Account) Print() {
	s := fmt.Sprintf("%s has a %s account with $%.2f",
		a.Name, a.AccountType.print(), a.Balance)
	fmt.Println(s)
}

// END PRINT OMIT

// START WITHDRAW OMIT

func (a *Account) Withdraw(amount float64) {
	a.Balance -= amount
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

// END WITHDRAW OMIT

func main() {
	a := Account{
		Name:        "Strans",
		Balance:     22.50,
		AccountType: Checking,
	}
	// START PRINT MAIN OMIT
	a.Print()
	// END PRINT MAIN OMIT

	// START WITHDRAW MAIN OMIT
	a.Withdraw(10.50)
	a.Print()

	a.Deposit(500)
	a.Print()
	// END WITHDRAW MAIN OMIT
}

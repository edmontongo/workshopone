package main

import (
	"math"
)

// DEPOSIT_S OMIT
func deposit(bank map[string]int32, name string, amount int32) bool {
	if !validateDeposit(bank, name, amount) {
		return false
	}
	bank[name] += amount
	return true
}

func validateDeposit(bank map[string]int32, name string, amount int32) bool {
	return bank != nil && amount >= 0 && math.MaxInt32-amount > bank[name]
}

// DEPOSIT_E OMIT

// WITHDRAW_S OMIT
func withdraw(bank map[string]int32, name string, amount int32, maxOverdraft int32) bool {
	if !validatewithdrawal(bank, name, amount, maxOverdraft) {
		return false
	}
	bank[name] -= amount
	return true
}

func validatewithdrawal(bank map[string]int32, name string, amount int32, maxOverdraft int32) bool {
	_, ok := bank[name]
	return ok && amount >= 0 && math.MinInt32+amount < bank[name] &&
		bank[name]-amount >= maxOverdraft
}

// WITHDRAW_E OMIT

func main() {

}

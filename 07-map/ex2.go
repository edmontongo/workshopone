package main

import (
	"fmt"
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
func withdraw(bank map[string]int, name string, amount int, maxOverdraft int) bool {
	if !validatewithdrawal(bank, name, amount, maxOverdraft) {
		return false
	}
	bank[name] -= amount
	return true
}

func validatewithdrawal(bank map[string]int, name string, amount int, maxOverdraft int) bool {
	_, ok := bank[name]
	return ok && amount >= 0 && math.MinInt32+amount < bank[name] &&
		bank[name]-amount >= maxOverdraft
}

// WITHDRAW_E OMIT

func main() {
	TestDeposit()
}

func TestDeposit() {
	pass := true
	for _, test := range depositTestData {
		bank := map[string]int32{}
		for k, v := range test.inbank {
			bank[k] = v
		}
		ret := deposit(bank, test.name, test.amount)
		if ret == test.ret && len(bank) == len(test.outbank) {
			for k, v := range test.outbank {
				if bank[k] != v {
					pass = false
				}
			}
		} else {
			pass = false
		}
		if !pass {
			fmt.Println("Fail withdraw(", test.inbank, ", \"", test.name, "\", ", test.amount, ")")
			fmt.Println(" expected ", test.ret, ", ", test.outbank)
			fmt.Println(" got ", ret, ",", bank)
			break
		}
		fmt.Println("Pass withdraw(", test.inbank, ", \"", test.name, "\", ", test.amount, ")")
		fmt.Println(" == ", test.ret, ", ", test.outbank)
	}
}

var depositTestData = []struct {
	inbank  map[string]int32
	name    string
	amount  int32
	ret     bool
	outbank map[string]int32
}{
	{
		map[string]int32{"Test": 100}, "Test", 100, true, map[string]int32{"Test": 200},
	},
	{
		map[string]int32{"Test": 100}, "Test", 100, true, map[string]int32{"Test": 201},
	},
}

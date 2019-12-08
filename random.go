package main

import (
	"math/rand"
	"time"
)

func RandomBank(bankCode int) BankAccount {
	var bank Bank
	if bankCode > 0 {
		bank = Banks[bankCode]
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
		ind := rand.Intn(len(Banks) - 1)
		count := 0
		for _, v := range Banks {
			if count == ind {
				bank = v
				break
			}
			count++
		}
	}

	first := bank.FirstBranch()

	rand.Seed(time.Now().UTC().UnixNano())
	branch := rand.Intn(first.Max - first.Min) + first.Min

	rand.Seed(time.Now().UTC().UnixNano())
	account := rand.Intn(99999999 - 990000) + 990000

	suffixes := []int{0, 2, 3, 30, 40, 50, 81}
	suffix := suffixes[rand.Intn(len(suffixes))]

	// No complicated maths here, just brute force to add one to account until it passes the modulus check
	for moduloCheck(bankCode, branch, account, suffix) != 0 {
		account++
	}

	b := BankAccount{}
	b.SetBank(bank.Code)
	b.SetBranch(branch)
	b.SetAccount(account)
	b.SetSuffix(suffix)

	return b
}

func moduloCheck(bank int, branch int, account int, suffix int) int {
	b := new(BankAccount)
	b.SetBank(bank)
	b.SetBranch(branch)
	b.SetAccount(account)
	b.SetSuffix(suffix)

	mod := BankModulo(bank)
	var cumulative int

	for i, n := range CheckDigitWeighting(bank, account) {
		cumulative += n * b.AccountNumber[i]
	}

	return cumulative % mod
}
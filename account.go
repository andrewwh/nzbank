package main

import (
	"fmt"
	"math"
	"errors"
)

type BankAccount struct {
	AccountNumber []int
}
func NewBankAccount(number string) (*BankAccount, error) {

	if len(number) != 16 {
		return nil, errors.New("Bank account must be 16 digits")
	}

	var digits = make([]int, 16)
	for i, v := range []byte(number[0: 16]) {
		if v >= '0' && v <= '9' {
			digits[i] = int(v - '0')
		}
	}

	acct := make([]int, 0)

	acct = append(acct, digits[0:6]...)
	acct = append(acct, 0)
	acct = append(acct, digits[6:13]...)
	acct = append(acct, 0)
	acct = append(acct, digits[13:]...)

	return &BankAccount{
			AccountNumber: acct,
		}, nil
}
func (b *BankAccount) SetBank(bank int) {
	for i, v := range IntToArray(bank, make([]int, 2)) {
		b.AccountNumber[i] = v
	}
}
func (b *BankAccount) SetBranch(branch int) {
	for i, v := range IntToArray(branch, make([]int, 4)) {
		b.AccountNumber[i+2] = v
	}
}
func (b *BankAccount) SetAccount(account int) {
	for i, v := range IntToArray(account, make([]int, 7)) {
		b.AccountNumber[i+7] = v
	}
}
func (b *BankAccount) SetSuffix(suffix int) {
	for i, v := range IntToArray(suffix, make([]int, 4)) {
		b.AccountNumber[i+14] = v
	}
}
func (b *BankAccount) GetBank() int {
	var bank int
	for i, n := range b.AccountNumber[0: 2] {
		bank += int(math.Pow(float64(10), float64(1-i))) * n
	}
	return bank
}
func (b *BankAccount) GetBranch() int {
	var branch int
	for i, n := range b.AccountNumber[2: 6] {
		branch += int(math.Pow(float64(10), float64(3-i))) * n
	}
	return branch
}
func (b *BankAccount) GetSuffix() int {
	var suffix int
	for i, n := range b.AccountNumber[14: 18] {
		suffix += int(math.Pow(float64(10), float64(3-i))) * n
	}
	return suffix
}
func (b *BankAccount) GetAccount() int {
	var account int
	for i, n := range b.AccountNumber[7: 14] {
		account += int(math.Pow(float64(10), float64(6-i))) * n
	}
	return account
}
func (b *BankAccount) IsValid() bool {
	if !b.IsValidBranch() {
		return false
	}

	mod := BankModulo(b.GetBank())
	var cumulative int

	for i, n := range CheckDigitWeighting(b.GetBank(), b.GetAccount()) {
		cumulative += n * b.AccountNumber[i]
	}

	return cumulative % mod == 0
}
func (b *BankAccount) IsValidBranch() bool {
	bank, found := Banks[b.GetBank()]

	if !found {
		return false
	}

	branch := b.GetBranch()
	for _, br := range bank.Branches {
		if br.Min <= branch && br.Max >= branch {
			return true
		}
	}

	return false
}
func (b *BankAccount) Print(formatted bool) string {
	if formatted {
		return fmt.Sprintf("%02d-%04d-%07d-%03d", b.GetBank(), b.GetBranch(), b.GetAccount(), b.GetSuffix())
	}

	return fmt.Sprintf("%02d%04d%07d%03d", b.GetBank(), b.GetBranch(), b.GetAccount(), b.GetSuffix())
}
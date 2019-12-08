package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	var validate string
	var bank int
	var random bool
	var formatted bool
	var silent bool

	flag.StringVar(&validate, "validate", "?", "Validate the give bank account")
	flag.BoolVar(&silent, "silent", false, "Do not output to stdout")
	flag.IntVar(&bank, "bank", 0, "Use bank during random generation")
	flag.BoolVar(&random, "random", false, "Randomly generate a bank account")
	flag.BoolVar(&formatted, "format", false, "Print in formatted form")
	flag.Parse()

	if validate != "?" {
		account := new(BankAccount)
		if err := account.Init(validate); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		isValid := account.IsValid()

		if isValid {
			if !silent {
				fmt.Printf("New Zealand Bank account %s is valid\n", account.Print(formatted))
			}
			os.Exit(0)
			
		} else {
			if !silent {
				fmt.Printf("New Zealand Bank account %s is not valid\n", account.Print(formatted))
			}
			os.Exit(1)
		}
	}

	if random {
		account := RandomBank(bank)
		fmt.Println(account.Print(formatted))
	}

}
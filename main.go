package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	var validate string
	var bank string
	var silent bool

	flag.StringVar(&validate, "validate", "?", "Validate the give bank account")
	flag.BoolVar(&silent, "silent", false, "Do not output to stdout")
	flag.StringVar(&bank, "random", "any", "Randomly generate a bank account for a bank")
	flag.Parse()

	if validate != "?" {
		account := new(BankAccount)
		account.Init(validate)
		isValid := account.IsValid()

		if isValid {
			if !silent {
				fmt.Printf("New Zealand Bank account %s is valid\n", account.ToString())
			}
			os.Exit(0)
			
		} else {
			if !silent {
				fmt.Printf("New Zealand Bank account %s is not valid\n", account.ToString())
			}
			os.Exit(1)
		}
	}
}
package main

import (
	"fmt"
	"scrapper/accounts"
)

func main() {
	account := accounts.Create("haril").Deposit(100)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}

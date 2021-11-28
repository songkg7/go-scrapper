package main

import (
	"fmt"
	"scrapper/accounts"
)

func main() {
	account := accounts.Create("haril")
	account.Deposit(10)
	fmt.Println(account.Balance())
}

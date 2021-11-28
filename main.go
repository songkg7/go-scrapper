package main

import (
	"fmt"
	"log"
	"scrapper/accounts"
)

func main() {
	account := accounts.Create("haril").Deposit(10)
	fmt.Println(account.Owner())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}

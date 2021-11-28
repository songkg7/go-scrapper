package main

import (
	"fmt"
	"scrapper/accounts"
)

func main() {
	account := accounts.Create("haril")
	fmt.Println(account)
}

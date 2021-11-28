package main

import "fmt"

func main() {
	fmt.Println(canIDrink(16))
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

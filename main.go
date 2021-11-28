package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	haril := person{
		name:    "haril",
		age:     28,
		favFood: []string{"ramen", "kimchi"},
	}
	fmt.Println(haril)
	fmt.Println(haril.name)
}

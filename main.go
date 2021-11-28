package main

import (
	"fmt"
	"scrapper/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "first word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
}

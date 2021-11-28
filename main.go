package main

import (
	"fmt"
	"scrapper/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "first word"}
	baseWord := "hello"
	errAdd := dictionary.Add(baseWord, "First")
	if errAdd != nil {
		fmt.Println(errAdd)
	}

	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}

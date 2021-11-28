package main

import (
	"fmt"
	"scrapper/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	totalLength, upper := lenAndUpper("haril")
	fmt.Println(totalLength, upper)
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

package main

import "fmt"

func main() {
	// Go 의 slice 는 immutable 하다.
	names := []string{"java", "python", "go"}
	newNames := append(names, "scala")
	fmt.Println(newNames)
}

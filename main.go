package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(multiply(2, 2))
	// multi value 는 반드시 값을 받아야하기 때문에 무시하기 위해서는 underscore 를 사용한다.
	totalLength, _ := lenAndUpper("haril")
	fmt.Println(totalLength)

	repeatMe("scala", "java", "python")
}

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

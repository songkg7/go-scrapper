package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	people := [2]string{"haril", "alex"}
	for _, person := range people {
		go isSexy(person, channel)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}
}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 10)
	channel <- person + " is sexy"
}

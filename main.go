package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)
	people := [2]string{"haril", "alex"}
	for _, person := range people {
		go isSexy(person, channel)
	}
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func isSexy(person string, channel chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	channel <- true
}

package main

import (
	"fmt"
	"time"
)


func SpaceOne(ch chan string) {
	ch <- "You are in space 1"
}

func SpaceTwo(ch chan string) {
	ch <- "You are in space 2"
}


func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go SpaceOne(ch1)
	go SpaceTwo(ch2)

	time.Sleep(1 * time.Second)
	select {
	case one := <-ch1:
		fmt.Println(one)
	case two := <- ch2:
		fmt.Println(two)
	default:
		fmt.Println("No space available")
	}

	
}

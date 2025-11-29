package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1: Buffered Channels
	fmt.Println("---")
	// Create a buffered channel with a capacity of 2.
	ch := make(chan int, 2)

	// Send two values without blocking.
	ch <- 1
	ch <- 2

	// Receive the values.
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Example 2: Select Statement
	fmt.Println("\n---")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// Example 3: Closing Channels
	fmt.Println("\n---")
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

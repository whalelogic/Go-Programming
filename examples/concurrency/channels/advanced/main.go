package main

import (
	"fmt"
	"time"
	"sync"
)

func Receiver(id int, ch <- chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range(ch) {
		fmt.Printf("Received ID: %d\t Value: %d\n", id, n)
	}
	fmt.Println("Exiting, done.")
}


func main() {
	// Example 1: Buffered Channels
	fmt.Println("---")
	ch := make(chan int, 2)

	// Send two values without blocking.
	ch <- 1
	ch <- 2

	// Receive the values.
	fmt.Println(<-ch)
	fmt.Println(<-ch)

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

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

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

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

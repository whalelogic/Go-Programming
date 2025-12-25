package main

import (
	"fmt"
	"os"
)

func printer(in <-chan int) {
	for x := range(in) {
		fmt.Println(x)
	}
}

func stringPrinter(in <-chan string) {
	for s := range(in) {
		fmt.Println(s)
	}
}

func main() {

	os.Setenv("USER", "gopher")
	fmt.Println("Hello, " + os.Getenv("USER"))
	os.Clearenv()

	fmt.Println("After clearing, USER is:", os.Getenv("USER"))
	fmt.Println("Nothing. It isn't set anymore.")

	c := make(chan int)
	go printer(c)

	i := 0
	for i < 9 {
		c <- i
		i++
	}
	close(c)

	// Dynamically modifying a slice and sending to a channel
	// The power of channels and goroutines
	var list = []string{"apple", "banana", "cherry"}
	f := make(chan string)
	go stringPrinter(f)
	for i := range list {
		f <- list[i] + " pie"
	}
	close(f)




}

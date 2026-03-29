package main

import (
	"fmt"
	"sync"
)

// chan<- string indicates channel is only for sending.

func SendChanMsg(msg string, ch chan<- string) {
	// SendChanMsg sends a message to a channel.

	// Start a new goroutine to send the message.
	go func() {
		ch <- msg
	}()
}

// Receiver --- ch <- chan means receiving
// Prints value as channels are received 
func Receiver(id int, ch <- chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range(ch) {
		fmt.Printf("Received ID: %d\t Value: %d\n", id, n)
	}
	fmt.Println("Exiting, done.")
}


func main() {
	// Create a new channel of type string.
	ch := make(chan string)

	// The message to send.
	s := "📮MESSAGE: Hello, Channel!"

	// Call the function to send the message to the channel.
	SendChanMsg(s, ch)

	// Wait and receive the message from the channel.
	receivedMsg := <-ch

	// Print the received message.
	fmt.Println("Received:", receivedMsg)
}

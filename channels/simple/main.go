package main

import "fmt"

// chan<- string indicates channel is only for sending.

func SendChanMsg(msg string, ch chan<- string) {
	// SendChanMsg sends a message to a channel.

	// Start a new goroutine to send the message.
	go func() {
		ch <- msg
	}()
}

func main() {
	// Create a new channel of type string.
	ch := make(chan string)

	// The message to send.
	s := "📮MESSAGE: from channel 📠) "
	// Call the function to send the message to the channel.
	SendChanMsg(s, ch)

	// Wait and receive the message from the channel.
	receivedMsg := <-ch

	// Print the received message.
	fmt.Println("Received:", receivedMsg)
}

package main

import (
	"fmt"
	"github.com/whalelogic/Go-Programming/stack"
)



// Create a stack using the stack package and perform some operations on it.
// s is the variable that holds this stack instance. It means create a Stack struct 
// from the stack package and assign it to s.
var s stack.Stack

// Type of s is stack.Stack, which is a struct defined in the stack package.
// View the methods of the stack package to see what operations we can perform on s.
// Completion should show methods like Push, Pop, Peek, Print, etc.
func main() {
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(13)
	s.Push(36)
	s.Print()
	
	output, b := s.Pop()
	s.Peek()
	s.Print()
	fmt.Println("Popped value:", output)
	fmt.Println("\nWas pop successful?", b)
	o := s.Peek()
	fmt.Println("\nPeek result:", o)
}

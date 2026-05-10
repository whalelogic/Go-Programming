package main

import (
	"fmt"
	"github.com/whalelogic/Go-Programming/data_structures/stack"
)



// Create a stack using the stack package and perform some operations on it.
// s is the variable that holds this stack instance. 

var s stack.Stack

// Type of s is stack.Stack, which is a struct defined in the stack package.
// Use methods like Push, Pop, Peek, and Print to manipulate the stack and display its contents.

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

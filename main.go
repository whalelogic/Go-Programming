package main

import (
	"fmt"

	"github.com/whalelogic/Go-Programming/stack"
	"github.com/whalelogic/Go-Programming/utils"
)



func main() {

	var s = "hello"
	fmt.Println(string(s[0]), string(s[1]), string(s[2]), string(s[3]), string(s[4]))
	fmt.Println(byte(s[0]))
	
	// Using the Swap function from the utils package
	x, y := utils.Swap("hello", 42)
	fmt.Println(x, y) 

	// Use the ParseURL function
	var URL = "https://www.example.com/path?query=123"
	parsed, err := utils.ParseURL(URL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	} else {
		fmt.Println("\nSuccessfully parsed URL: \t", parsed)
	}

	// Create a stack 
	type Node struct {
	data int 
	next *Node
	}

	type Stack struct {
	top *Node
	length int
	}

	arr := stack.Stack{}
	arr.Push(10)
	arr.Push(20)
	arr.Push(30)
	arr.Push(40)
	arr.Print()
	arr.Peek()
	arr.Pop()
	arr.Print()
	arr.Peek()

	fmt.Println(arr)








}

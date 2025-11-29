package stack

import (
	"fmt"
)


type Node struct {
	data int 
	next *Node
}

type Stack struct {
	top *Node
	length int
}

func(s *Stack) Push(data int) {
	newNode := &Node {
		data: data,
		next: s.top,
	}
	s.top = newNode
	s.length++
	fmt.Println(newNode.data)
}

func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	val := s.top.data 
	s.top = s.top.next // move item from top - down 1 node.
	s.length--
	fmt.Printf("Top of the stack is: %#v\n", s.top.data)
	return val, true
}


func (s *Stack) Peek() (bool) {
	if s.top == nil {
		fmt.Println("Stack is empty!")
	} else if s.top != nil {
		fmt.Printf("Stack has items: %#v\t length: %v", s.top, s.length)
	}
	return true
}


// function to print the stack 

func (s *Stack) Print() {
	current := s.top 
	fmt.Println("Stack top -> bottom: ")
	for current != nil {
		fmt.Printf("Current stack: %v\n", current.data)
		current = current.next
	}
	fmt.Println()
}


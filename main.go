package main

import (
	"fmt"
	"strconv"

	"github.com/whalelogic/Go-Programming/lists"
	"github.com/whalelogic/Go-Programming/sort"
	"github.com/whalelogic/Go-Programming/stack"
)


var myList lists.List[int]

func main() {
	// Stack example
	// Using our sort functions on it
	arr2 := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	sortedArr := sort.Sort(arr2)
	fmt.Println(sortedArr)
	arr := stack.Stack{}
	arr.Push(10)
	arr.Print()

	// Linked List example
	myList.PushBack(10)
	myList.PushBack(20)
	myList.PushFront(5)
	myList.InsertAfter(myList.Front(), 7)
	myList.InsertAfter(myList.Front(), 13)

	fmt.Println("Linked List elements:")
	// loop through the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// Remember, Linked Lists are structs
	// We access their fields using dot notation

	// Built-in methods cap won't work here
	// fmt.Println(len(myList)) <-- Won't work either
	// Hence the need for a Length() method (Declared in lists/linkedList.go)
	fmt.Println("First element: ", myList.Front().Value)
	fmt.Println("Last element: ", myList.Back().Value)
	fmt.Println("Length of linked-list: ", myList.Length())
	sl := myList.MakeSlice()
	for i := range sl {
		// Must convert int to string to concatenate
		fmt.Println(strconv.Itoa(i) + ": " + strconv.Itoa(sl[i]))
	}
	
	v, ok := myList.RemoveFront()
	if ok {
		fmt.Println("Removed front element: ", v)
	}
	var myList2 lists.List[string]
	myList2.PushBack("Hello")
	myList2.PushBack("World")
	for e := myList2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("First element in second list", myList2.Front().Value)



}

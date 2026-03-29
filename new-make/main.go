package main

import (
	"fmt"
)


func main() {
	


	// New and make are both used to allocate memory, but they serve different purposes.
	var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
	var v  []int = make([]int, 10) // the slice v now refers to a new array of 100 ints
	fmt.Println(v)

	v = append(v, 1, 2, 3)
	fmt.Println(v)
	fmt.Println("p: ", p)
	fmt.Println("v: ", v)

	fmt.Printf("p: %v\n", p)



}




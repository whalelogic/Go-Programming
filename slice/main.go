package main 

import (
	"fmt"
	"slices"
)

func main() {

	n := []int{1,3,5,8,11,13,24,71,54,42,103}
	copy := slices.Clone(n)

	// nums in n
	for _, n := range n {
		fmt.Println("N =", n)
	}

	for o, n := range n {
		fmt.Printf("Index: %v --> Value: %v\n", o, n)
	}

	// replicate the slice with built-in 'slices.Clone'
	fmt.Println(copy)





	

}

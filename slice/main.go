package main

import (
	"fmt"
	"slices"
)

func main() {

	n := []int{1,3,5,8,11,13,24,71,54,42,103}
	copy := slices.Clone(n)

	double := func(x float64) float64 {
		return x * 2
	}

	fmt.Printf("Double 4.5=%+v\n", double(4.5))

	// every index in n...🢙
	for o := range n {
		fmt.Println("Index:", o)
		fmt.Println("Index x 2 = ", o*2)
	}

	// every index and int in n...🢙
	for o, n := range n {
		fmt.Printf("Index: %v --> Value: %v\n", o, n)
	}

	fmt.Println(copy)

	slice2 := slices.Delete(copy, 2,5)
	fmt.Println("After Deletion:", slice2)
	both := slices.Contains(n, 24)
	fmt.Println("Contains 24?:", both)

	combine := slices.Concat(n, slice2)
	fmt.Println("Combine both slices: ", combine)

	// Now sort the combined slice 
	slices.Sort(combine)
	fmt.Println("Sorted combined slice: ", combine)

	// Now print even and odd values from the combined slice
	for n, v := range combine {
		if v%2 == 0 {
			fmt.Printf("Index: %v --> Even Value: %v\n", n, v)
		} else {
			fmt.Printf("Index: %v --> Odd Value: %v\n", n, v)
		}
	}

	// Now we'll remove the duplicates from the combined slice
	uniqueSlice := slices.Compact(combine)
	fmt.Println("Unique values from combined slice: ", uniqueSlice)

	s, b := slices.BinarySearch(combine, 103)
	fmt.Printf("Position: %d  Found: %v\n", s, b)
	// False??
	// Because slices.BinarySearch() requires the 
	// slice be pre-sorted in increasing order
	slices.Sort(combine)
	s, b = slices.BinarySearch(combine, 103)
	fmt.Printf("\nIs 103 in the slice? Where?\nFound: %v --> Position: %d\n", b, s)
	fmt.Println("Now it works!")



	}



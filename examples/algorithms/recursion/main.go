package main

import (
	"fmt"
	"time"
)


func Sum(n int) int {
	if n <= 0 {
		return 0
	}
	return n + Sum(n-1)
}

func SumSlice(n []int) int {
	t := time.Now()
	sum := 0
	for _, val := range n {
		sum += val
	}
	fmt.Println("SumSlice speed: ", time.Since(t))
	return sum
}

func RecursiveSum(n []int) (int) {
	// This len(n) control statement is necessary or it will run forever or not at all.
	if len(n) == 0 {
		return 0
	}
	return n[0] + RecursiveSum(n[1:])
}

func CheckEvenNumbers(nums []int) (int, bool) {
	var b = false
	total := 0
	if len(nums) <= 1 {
		return 0, false
	}
	for i, n := range nums {
		if n % 2 == 0 {
			b = true
			total++
			return i, b
		}

	}
	fmt.Println("Contains even numbers? ", b)
	fmt.Println("How many? ", total)
	return total, b

}


var list = []int{3,5,9,13,18, 44, 16, 4}


func main() {
	fmt.Println("Sum of n: ", Sum(3))

	fmt.Println("SumSlice function: ", SumSlice(list))

	fmt.Println("Recursive Sum function/speed: ", RecursiveSum(list))

	i, c := CheckEvenNumbers(list)
	
	// If 'c' is 'true', evens exist and 'c' contains the number of evens found.
	if c {
		fmt.Println(c)
		fmt.Println("Evens:", i)
	}

	
}

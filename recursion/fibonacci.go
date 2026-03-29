// Package recursion implementations from each package 
package recursion 

import "fmt"


// FibRecursive generates Fibonacci numbers recursively.
func FibRecursive(n int) int {
	if n <= 0 {
	return 0
	} else if n == 1 {
		return 1
	} else {
		return FibRecursive(n-1) + FibRecursive(n-2)
	} 
}


func Fibonacci(n int) (result int) {
	if n <= 1 {
		return n
	} else {
		result = Fibonacci(n-1) + Fibonacci(n-2)
	}
	return result
}


func GenerateFibs(series []int) []int {
		// Must use at least 2 integers
		for len(series) < 2 {
			fmt.Println("Series must have at least two integers")
			break
	}
		for range(series) {
			// Get the last element
			last := len(series) - 1
			// Get the second to last element
			// Add them together
			// Append to the series
			next := series[last] + series[last-1]
			// Append result to the series

			series = append(series, next)
}
		return series
	
}






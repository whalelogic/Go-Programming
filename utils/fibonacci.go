// Package utils implementations from each package 
package utils 


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






// Package sort provides basic sorting functionalities similar to 
// the Go standard library.
package sort

func Sort(n []int) []int {
	for i := range n {
		for j := i + 1; j < len(n); j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	return n
}

func SwapInts(a, b *int) (int, int) {
	*a, *b = *b, *a
	return *a, *b
}


func SwapIntsNoPtr(a, b int) (int, int) {
	return b, a
}

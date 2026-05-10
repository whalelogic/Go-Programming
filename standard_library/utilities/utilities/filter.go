package utilities

// FilterInPlace filters slice elements in place.
// Checks if vals are divisible by zero. 
func FilterInPlace(n []int) []int {
	i := 0	
	for _, v := range n {
		if v%2 == 0 {
			n[i] = v 
			i++
		}
	}
	return n[:i]
}



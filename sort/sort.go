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

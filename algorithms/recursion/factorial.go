package recursion

func Factor(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factor(n-1)
}

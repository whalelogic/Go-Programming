package utilities

func ReverseArray[T any](s []T) []T {
	n := len(s)
	out := make([]T, n)

	for i := range(n) {
		out[n-1-i] = s[i]
	}
	return out
}

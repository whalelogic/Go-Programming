package utils

func ReverseArrayCopy[T any](s []T) []T {
	n := len(s)
	out := make([]T, n)

	for i := 0; i < n; i++ {
		out[n-1-i] = s[i]
	}
	return out
}

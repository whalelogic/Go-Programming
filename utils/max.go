package utils

import "fmt"

func Max(num []int) int {
	var max int
	for _, n := range num {
		if n > max {
			max = n
			fmt.Println("New max:", max)
		}
	}
	return max
}

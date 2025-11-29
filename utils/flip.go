package utils

import (
	"fmt"
)

func Flip(n []int) []int {
	for i := 0; i < len(n); i++ {
		j := len(n) -1 -i
		fmt.Println("i: ", i, "j: ", j)
	}
	return n
}

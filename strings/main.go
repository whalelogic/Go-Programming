package main 

import (
	"fmt"
	"strings"
)


// ReverseString can be simplified. 
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	var question = "Why is the sky blue?"
	fmt.Println(question)

	// To capitalize the entire string use something like:
	fmt.Println(strings.ToTitle(question))
	var split = strings.Split(question, "")
	fmt.Println(split)

}

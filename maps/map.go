package main

import (
	"fmt"
	"strings"
)


func StringReader(prompt string) string {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println("Error reading input, please try again.")
		return StringReader(prompt)
	}
	return input
}

func StringTrim(input string) string {
	return strings.TrimSpace(input)
}

func GetInput(prompt string) string {
	fmt.Print(prompt)
	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		return input
	}
}


func PrintNestedMap(nestedMap map[string]map[string]string) {
	for key, innerMap := range nestedMap {
		fmt.Printf("Key: %s\n", key)
		for innerKey, value := range innerMap {
			fmt.Printf("  %s: %s\n", innerKey, value)
		}
	}
}


func main() {
	i := GetInput("Enter something: ")
	fmt.Println("You entered:", i)

	m := map[string]map[string]string{
		"outer1": {
			"inner1": i,
			"inner2": "value2",
		},
		"outer2": {
			"inner3": i,
			"inner4": "value4",
		},
	}

	PrintNestedMap(m)
}

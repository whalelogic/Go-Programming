package maps

import (
	"fmt"
	"os"
	"io"
)

func GetInput(prompt string) any {
	fmt.Print(prompt)
	var input any
	r := io.Reader(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return input
}



func NestedStringMap() map[string]map[string]string {
	nestedMap := make(map[string]map[string]string)
	fmt.Println("Enter Name and Country followed by enter: ")
	name, country := io.ReadString('\n'), io.ReadString('\n')
	if name == "" || country == "" {
		fmt.Println("Invalid input. Name and Country cannot be empty.")
		os.Exit(1)
	}
	nestedMap["Person1"] = map[string]string{
		"Name":    name,
		"Country": country,
	}
	return nestedMap
}

func PrintNestedMap(nestedMap map[string]map[string]string) {
	for key, innerMap := range nestedMap {
		fmt.Printf("Key: %s\n", key)
		for innerKey, value := range innerMap {
			fmt.Printf("  %s: %s\n", innerKey, value)
		}
	}
}

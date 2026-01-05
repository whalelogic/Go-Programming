package maps

import (
	"fmt"
)

func GetInput(prompt string, r string) any {
	fmt.Print(prompt)
	fmt.Scan(&r)
	return r
}

func PrintNestedMap(nestedMap map[string]map[string]string) {
	for key, innerMap := range nestedMap {
		fmt.Printf("Key: %s\n", key)
		for innerKey, value := range innerMap {
			fmt.Printf("  %s: %s\n", innerKey, value)
		}
	}
}

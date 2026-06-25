package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Go Built-in Functions Demo ---")

	// 1. make and len/cap
	slice := make([]int, 3, 5)
	slice[0], slice[1], slice[2] = 10, 20, 30
	fmt.Printf("Slice: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	// 2. append
	slice = append(slice, 40, 50, 60)
	fmt.Printf("After append: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	// 3. copy
	dest := make([]int, len(slice))
	numCopied := copy(dest, slice)
	fmt.Printf("Copied %d elements: %v\n", numCopied, dest)

	// 4. delete (maps)
	m := map[string]string{"A": "Apple", "B": "Banana", "C": "Cherry"}
	fmt.Println("Map before delete:", m)
	delete(m, "B")
	fmt.Println("Map after delete 'B':", m)

	// 5. complex, real, imag
	c := complex(3, 4)
	fmt.Printf("Complex: %v, Real: %v, Imag: %v\n", c, real(c), imag(c))

	// 6. new
	p := new(int)
	*p = 100
	fmt.Printf("New int pointer value: %d\n", *p)

	// 7. min, max (Go 1.21+)
	// Note: These might fail if the environment Go version is < 1.21
	// but are included for completeness as requested.
	fmt.Printf("Min(10, 20, 5): %d\n", min(10, 20, 5))
	fmt.Printf("Max(10, 20, 5): %d\n", max(10, 20, 5))

	// 8. clear (Go 1.21+)
	clear(m)
	fmt.Println("Map after clear:", m)
	clear(slice)
	fmt.Println("Slice after clear (zeroed):", slice)

	// 9. panic and recover
	demoPanic()
}

func demoPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Successfully recovered from panic:", r)
		}
	}()

	fmt.Println("Triggering a panic...")
	panic("wrOops!!!!!")
}

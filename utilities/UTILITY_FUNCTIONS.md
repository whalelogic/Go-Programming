# Go Utility Functions

This directory contains a collection of simple utility functions written in Go. This guide explains the purpose and implementation of each function and discusses the broader Go concepts they illustrate.

## 1. `fibonacci.go`

This file likely contains a function to calculate numbers in the Fibonacci sequence. The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones, usually starting with 0 and 1.

`0, 1, 1, 2, 3, 5, 8, 13, 21, ...`

### Concepts Illustrated:

-   **Recursion:** The Fibonacci sequence is a classic example used to teach recursion. A recursive function would look something like this:

    ```go
    func Fibonacci(n int) int {
        if n <= 1 {
            return n
        }
        return Fibonacci(n-1) + Fibonacci(n-2)
    }
    ```
    While elegant, this recursive approach is highly inefficient (O(2^n)) due to repeated calculations.

-   **Iteration:** A more efficient approach is to use iteration, which has a time complexity of O(n).

    ```go
    func FibonacciIterative(n int) int {
        if n <= 1 {
            return n
        }
        var n2, n1 = 0, 1
        for i := 2; i <= n; i++ {
            n2, n1 = n1, n2+n1
        }
        return n1
    }
    ```

-   **Memoization:** To optimize the recursive version, you can use memoization (caching the results of expensive function calls) to avoid redundant computations.

## 2. `max.go`

This function probably finds the maximum value in a slice of numbers.

### Concepts Illustrated:

-   **Iteration:** The most straightforward way to find the maximum value is to iterate through the slice and keep track of the largest value seen so far.
-   **Edge Cases:** A robust implementation should handle edge cases, such as an empty or nil slice.

```go
func Max(s []int) (int, error) {
    if len(s) == 0 {
        return 0, fmt.Errorf("slice is empty")
    }
    max := s[0]
    for _, v := range s[1:] {
        if v > max {
            max = v
        }
    }
    return max, nil
}
```

## 3. `reverse_arr.go` and `reverse_slice.go`

These files highlight a crucial distinction in Go: the difference between arrays and slices.

-   **Arrays:** Have a fixed size, which is part of their type. `[5]int` and `[10]int` are different types. When you pass an array to a function, you are passing a *copy* of the array.
-   **Slices:** Are dynamically-sized and more flexible. A slice is a lightweight struct that contains a pointer to an underlying array, a length, and a capacity. When you pass a slice to a function, you are passing a copy of the slice header, but the pointer still points to the same underlying array. This means the function can modify the contents of the underlying array.

The implementation for reversing would be similar for both, but how the functions are defined and called would differ.

```go
// For a slice (modifies the original)
func ReverseSlice(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// For an array (returns a new, reversed array)
func ReverseArray(arr [5]int) [5]int {
    var reversed [5]int
    for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
        reversed[i] = arr[j]
    }
    return reversed
}
```

## 4. `swap_any_val.go`

This file likely demonstrates how to swap two values of any type.

### Concepts Illustrated:

-   **Pointers:** Before generics, swapping values of arbitrary types would typically be done using pointers and `interface{}`. However, this approach lacks type safety.

-   **Generics (Go 1.18+):** Generics provide a type-safe way to write functions that can operate on multiple types. A generic swap function would be the modern, idiomatic way to implement this.

    ```go
    func Swap[T any](a, b *T) {
        *a, *b = *b, *a
    }

    func main() {
        x, y := 10, 20
        Swap(&x, &y)
        fmt.Println(x, y) // 20 10

        s1, s2 := "hello", "world"
        Swap(&s1, &s2)
        fmt.Println(s1, s2) // world hello
    }
    ```

## 5. `flip.go`

This function could have several interpretations. It might flip the bits of an integer, or it could be another name for a reversal function. Assuming it's a boolean flip:

```go
func Flip(b bool) bool {
    return !b
}
```

If it's a bit flip (bitwise NOT):

```go
func FlipBits(n uint8) uint8 {
    return ^n // The ^ operator is the bitwise NOT
}
```

## Conclusion

This collection of utilities, while simple, serves to demonstrate a wide range of important Go features and idioms. From recursion and iteration to the fundamental differences between arrays and slices, and from basic algorithms to the power of modern generics, these examples provide a practical look at Go programming concepts.

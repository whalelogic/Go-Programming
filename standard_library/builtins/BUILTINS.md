# Go Built-in Functions Reference

Go provides a set of built-in functions that are available in every Go program without importing any package. These functions provide core language features like slice/map manipulation, memory allocation, and concurrency primitives.

## Quick Reference Table

| Function | Signature | Description |
| :--- | :--- | :--- |
| **append** | `func append(s []T, vs ...T) []T` | Appends elements to a slice and returns the new slice. |
| **cap** | `func cap(v Type) int` | Returns the capacity of a slice, array, or channel. |
| **clear** | `func clear(t Type)` | Clears all elements from a map or sets all elements of a slice to zero values (Go 1.21+). |
| **close** | `func close(c chan<- T)` | Closes a channel. |
| **complex** | `func complex(r, i FloatType) ComplexType` | Creates a complex number from real and imaginary parts. |
| **copy** | `func copy(dst, src []T) int` | Copies elements from a source slice to a destination slice. |
| **delete** | `func delete(m map[K]V, k K)` | Deletes the element with the specified key from a map. |
| **imag** | `func imag(c ComplexType) FloatType` | Returns the imaginary part of a complex number. |
| **len** | `func len(v Type) int` | Returns the length of a string, slice, array, map, or channel. |
| **make** | `func make(t Type, size ...IntegerType) Type` | Initializes slices, maps, and channels. |
| **max** | `func max(x T, y ...T) T` | Returns the maximum value of a fixed number of arguments (Go 1.21+). |
| **min** | `func min(x T, y ...T) T` | Returns the minimum value of a fixed number of arguments (Go 1.21+). |
| **new** | `func new(Type) *Type` | Allocates zeroed memory for a type and returns a pointer to it. |
| **panic** | `func panic(v interface{})` | Stops normal execution of the current goroutine. |
| **print** | `func print(args ...Type)` | Formats arguments and writes them to standard error (primitive). |
| **println** | `func println(args ...Type)` | Like `print` but adds a newline. |
| **real** | `func real(c ComplexType) FloatType` | Returns the real part of a complex number. |
| **recover** | `func recover() interface{}` | Allows a program to manage behavior of a panicking goroutine. |

---

## Detailed Usage Examples

### 1. Memory Allocation: `make` vs `new`

`make` is used for initializing built-in reference types (slices, maps, channels), while `new` is used for allocating memory and returning a pointer to a zeroed value.

```go
// make(Type, length, capacity)
s := make([]int, 5, 10) 

// new(Type) returns *Type
p := new(int) 
*p = 42
```

### 2. Collection Manipulation: `append`, `copy`, `delete`, `clear`

```go
// Append
s := []int{1, 2}
s = append(s, 3, 4)

// Copy
dst := make([]int, len(s))
count := copy(dst, s)

// Delete from map
m := map[string]int{"a": 1, "b": 2}
delete(m, "a")

// Clear (Go 1.21+)
clear(m) // m is now empty
clear(s) // all elements of s are now 0
```

### 3. Comparison: `min`, `max` (Go 1.21+)

```go
smallest := min(10, 20, 5, 30) // 5
largest := max(10, 20, 5, 30)  // 30
```

### 4. Panics and Recovery

```go
func handlePanic() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
    }
}

func causePanic() {
    defer handlePanic()
    panic("something went wrong")
}
```

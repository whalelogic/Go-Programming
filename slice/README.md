# Go Slices: A Deep Dive

> Slice package contains a main.go file with examples of using slices in Go.

Run the code in `main.go` to see examples in action.

This guide provides a comprehensive look at slices in Go. Slices are one of the most powerful and versatile features of Go, and a solid understanding of them is crucial for any Go developer.

## 1. What is a Slice?

A slice is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

A slice is described by a struct with three fields:

-   **Pointer:** A pointer to the first element of the underlying array that is reachable through the slice.
-   **Length:** The number of elements in the slice.
-   **Capacity:** The maximum number of elements the slice can hold, starting from the pointer.

You can get the length and capacity of a slice using the built-in `len()` and `cap()` functions.

```go
s := make([]int, 5, 10)
fmt.Println("len:", len(s)) // 5
fmt.Println("cap:", cap(s)) // 10
```

## 2. Creating and Initializing Slices

There are several ways to create a slice:

### Using a Slice Literal

This is the most common way to create a slice.

```go
s := []int{1, 2, 3, 4, 5}
```

### Using `make()`

The `make()` function allows you to create a slice with a specified length and capacity.

```go
s := make([]int, 5)      // len=5, cap=5
s := make([]int, 5, 10)   // len=5, cap=10
```

### Slicing an Existing Array or Slice

You can create a new slice by "slicing" an existing array or slice.

```go
arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
s1 := arr[2:5]  // s1 = [2, 3, 4], len=3, cap=8
s2 := s1[1:3]   // s2 = [3, 4], len=2, cap=7
```

**Important:** When you slice an existing array or slice, the new slice shares the same underlying array. This means that if you modify an element in the new slice, the change will be reflected in the original array/slice.

```go
s1[1] = 99
fmt.Println(arr) // [0 1 2 99 4 5 6 7 8 9]
```

## 3. Common Slice Operations

### Appending to a Slice

The `append()` built-in function is used to add elements to the end of a slice.

```go
s := []int{1, 2, 3}
s = append(s, 4, 5) // s is now [1, 2, 3, 4, 5]
```

`append()` is smart about memory allocation. If the underlying array has enough capacity, the new elements are added to the existing array. If not, a new, larger array is allocated, the elements from the old slice are copied to the new array, and the new elements are added. The resulting slice will then point to this new array.

### Iterating over a Slice

The `for...range` loop is the idiomatic way to iterate over a slice.

```go
for index, value := range s {
    fmt.Printf("index: %d, value: %d\n", index, value)
}
```

If you don't need the index, you can ignore it with the blank identifier `_`.

```go
for _, value := range s {
    fmt.Println(value)
}
```

### Copying a Slice

The `copy()` built-in function copies elements from a source slice to a destination slice. It returns the number of elements copied.

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
numCopied := copy(dst, src)
```

`copy()` is useful when you want to create a new slice that is a true copy of the original, not just a view into the same underlying array.

## 4. The `slices` Package (Go 1.21+)

Go 1.21 introduced the new `slices` package, which provides a collection of generic functions that are useful with slices of any type. Some of the most useful functions include:

-   `slices.Sort(s)`: Sorts the elements of a slice.
-   `slices.Compact(s)`: Removes consecutive duplicate elements.
-   `slices.Delete(s, i, j)`: Deletes a range of elements.
-   `slices.Clone(s)`: Creates a shallow copy of a slice.
-   `slices.Contains(s, v)`: Checks if a value is present in a slice.
-   `slices.Index(s, v)`: Returns the index of the first occurrence of a value.

The `main.go` file in this directory provides examples of using some of these functions.

## 5. Slice Gotchas

-   **Nil vs. Empty Slice:** A nil slice has a length and capacity of 0, but its pointer is `nil`. An empty slice also has a length and capacity of 0, but its pointer is not `nil`. For most practical purposes, you can treat them the same.
-   **Appending to a Sliced Slice:** Be careful when appending to a slice that was created by slicing another slice. If the original slice had extra capacity, the append operation might overwrite values in the original slice that are beyond the new slice's length but within its capacity.

## Conclusion

Slices are a fundamental and powerful part of Go. By understanding their internal structure and how they relate to arrays, you can use them effectively and avoid common pitfalls. The new `slices` package adds even more power and convenience, making slice manipulation easier and more readable.

# Sorting in Go

This guide covers how to sort data in Go using the standard library's `sort` package.

## Quick Reference: `sort` Package

| Function / Method | Signature | Description |
| :--- | :--- | :--- |
| **Ints** | `func Ints(x []int)` | Sorts a slice of ints in ascending order. |
| **Strings** | `func Strings(x []string)` | Sorts a slice of strings in ascending order. |
| **Float64s** | `func Float64s(x []float64)` | Sorts a slice of float64s in ascending order. |
| **Slice** | `func Slice(x any, less func(i, j int) bool)` | Sorts a slice using a custom `less` function. |
| **SliceStable**| `func SliceStable(x any, less func(i, j int) bool)` | Performs a stable sort (keeps order of equals). |
| **Sort** | `func Sort(data Interface)` | Sorts a collection implementing `sort.Interface`. |
| **SearchInts** | `func SearchInts(a []int, x int) int` | Binary search for `x` in sorted slice `a`. |
| **Search** | `func Search(n int, f func(int) bool) int` | Generic binary search for index `i` where `f(i)` is true. |
| **IsSorted** | `func IsSorted(data Interface) bool` | Checks if a collection is sorted. |
| **Reverse** | `func Reverse(data Interface) Interface` | Returns a reversed view of the collection. |

---

## 1. The `sort` Package

The `sort` package provides primitives for sorting slices and user-defined collections. It offers several ways to sort data, from simple cases with basic types to complex sorting of custom data structures.

## 2. Sorting Basic Types

For the basic types (`int`, `float64`, `string`), the `sort` package provides convenient functions:

-   `sort.Ints(s []int)`
-   `sort.Float64s(s []float64)`
-   `sort.Strings(s []string)`

These functions sort the provided slice in place (in ascending order).

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    ints := []int{4, 2, 8, 1, 5}
    sort.Ints(ints)
    fmt.Println("Sorted ints:", ints) // [1 2 4 5 8]

    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Sorted strs:", strs) // [a b c]
}
```

To check if a slice is already sorted, you can use:

-   `sort.IntsAreSorted(s []int)`
-   `sort.Float64sAreSorted(s []float64)`
-   `sort.StringsAreSorted(s []string)`

## 3. The `sort.Interface`

For custom data types, you need to implement the `sort.Interface`. This interface has three methods:

-   **`Len() int`**: Returns the number of elements in the collection.
-   **`Less(i, j int) bool`**: Reports whether the element with index `i` should sort before the element with index `j`.
-   **`Swap(i, j int)`**: Swaps the elements with indexes `i` and `j`.

Once your type implements this interface, you can use `sort.Sort()` to sort a slice of that type.

### Example: Sorting a Slice of Structs

Let's say we have a slice of `Person` structs and we want to sort them by age.

```go
type Person struct {
    Name string
    Age  int
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
    people := []Person{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }

    sort.Sort(ByAge(people))

    fmt.Println(people) // [{Michael 17} {Jenny 26} {Bob 31} {John 42}]
}
```

This pattern of creating a new type (e.g., `ByAge`) that is a slice of your custom type is a common and idiomatic way to handle sorting in Go.

## 4. Using `sort.Slice`

While `sort.Interface` is powerful, it can be a bit verbose. Go 1.8 introduced `sort.Slice`, which provides a more convenient way to sort slices without needing to implement `sort.Interface`.

`sort.Slice` takes a slice and a `less` function as arguments. The `less` function is a closure that captures the slice being sorted.

```go
func main() {
    people := []Person{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }

    // Sort by age
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })
    fmt.Println(people)

    // Sort by name
    sort.Slice(people, func(i, j int) bool {
        return people[i].Name < people[j].Name
    })
    fmt.Println(people)
}
```

`sort.Slice` is often preferred for its conciseness and flexibility, as you can easily define different sorting orders without creating new types.

There is also `sort.SliceStable`, which performs a stable sort. A stable sort keeps the original order of equal elements.

## 5. Searching in a Sorted Slice

The `sort` package also includes functions for searching in a sorted slice, which is much more efficient (O(log n)) than a linear scan (O(n)).

`sort.Search()` performs a binary search. It takes the length of the slice and a "search" function that returns `true` if the element at a given index is "at or after" the desired value.

```go
ints := []int{1, 2, 4, 5, 8}
target := 4

index := sort.Search(len(ints), func(i int) bool {
    return ints[i] >= target
})

if index < len(ints) && ints[index] == target {
    fmt.Printf("found %d at index %d\n", target, index)
} else {
    fmt.Printf("%d not found\n", target)
}
```

For slices of basic types, you can use the more convenient `sort.SearchInts()`, `sort.SearchFloat64s()`, and `sort.SearchStrings()`.

## Conclusion

The `sort` package in Go provides a comprehensive set of tools for sorting data. Whether you're sorting simple slices of built-in types or complex custom data structures, Go has a solution that is efficient, flexible, and idiomatic. For most custom sorting needs, `sort.Slice` is the recommended approach due to its clarity and conciseness.

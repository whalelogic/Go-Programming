# Go Maps: Hash Tables in Go

Maps are Go's built-in associative data type (sometimes called hashes or dicts in other languages). They provide a way to store data as key-value pairs.

## Quick Reference: Map Operations

| Operation | Syntax / Function | Description |
| :--- | :--- | :--- |
| **Declaration** | `var m map[K]V` | Declares a map (initially nil). |
| **Initialization** | `make(map[K]V, hint)` | Initializes a map with an optional capacity hint. |
| **Literal** | `m := map[K]V{k1: v1}` | Initializes a map with values. |
| **Insertion** | `m[key] = value` | Adds or updates an element. |
| **Retrieval** | `val := m[key]` | Retrieves a value (returns zero value if key missing). |
| **Check Existence**| `val, ok := m[key]` | Retrieves value and a boolean indicating if key exists. |
| **Deletion** | `delete(m, key)` | Removes a key-value pair. |
| **Length** | `len(m)` | Returns the number of key-value pairs. |
| **Clear** | `clear(m)` | Removes all keys from the map (Go 1.21+). |
| **Iteration** | `for k, v := range m` | Iterates over all key-value pairs (order is random). |

---

## 1. Map Basics

### Initialization
```go
// Using make
scores := make(map[string]int)

// Using a literal
colors := map[string]string{
    "red":   "#ff0000",
    "green": "#00ff00",
}
```

### Accessing Values
```go
val, ok := scores["Alice"]
if ok {
    fmt.Println("Alice's score:", val)
} else {
    fmt.Println("Alice not found")
}
```

## 2. Important Characteristics

1.  **Unordered:** Iteration order over maps is not guaranteed and can change between executions.
2.  **Reference Type:** Like slices, maps are reference types. Passing a map to a function allows the function to modify the original map.
3.  **Key Constraints:** Map keys can be of any type for which the equality operator `==` is defined (e.g., integers, floats, strings, pointers, channels, interfaces, and structs that only contain comparable types). Slices, maps, and functions cannot be map keys.

## 3. The `maps` Package (Go 1.21+)

The standard library now includes a `maps` package for common map operations:

-   `maps.Keys(m)`: (Proposed/Planned) Returns a slice of map keys.
-   `maps.Values(m)`: (Proposed/Planned) Returns a slice of map values.
-   `maps.Copy(dst, src)`: Copies all key/value pairs from src to dst.
-   `maps.Equal(m1, m2)`: Reports whether two maps contain the same key/value pairs.
-   `maps.DeleteFunc(m, f)`: Deletes any key/value pairs from m for which f(k, v) returns true.

## Example Code

See `data_structures/maps/map.go` for more examples.

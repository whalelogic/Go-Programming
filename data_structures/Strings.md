# String Manipulation in Go

This guide provides a comprehensive overview of working with strings in Go.

## Quick Reference: Strings & Conversion

| Package | Function | Signature | Description |
| :--- | :--- | :--- | :--- |
| **strings** | `Contains` | `func Contains(s, substr string) bool` | Checks if `substr` is in `s`. |
| **strings** | `HasPrefix` | `func HasPrefix(s, prefix string) bool` | Checks if `s` starts with `prefix`. |
| **strings** | `HasSuffix` | `func HasSuffix(s, suffix string) bool` | Checks if `s` ends with `suffix`. |
| **strings** | `Index` | `func Index(s, substr string) int` | Returns index of first `substr`. |
| **strings** | `Join` | `func Join(elems []string, sep string) string` | Joins strings with a separator. |
| **strings** | `Split` | `func Split(s, sep string) []string` | Splits string into a slice. |
| **strings** | `ReplaceAll`| `func ReplaceAll(s, old, new string) string` | Replaces all instances of `old`. |
| **strings** | `TrimSpace` | `func TrimSpace(s string) string` | Removes leading/trailing whitespace. |
| **strings** | `ToUpper` | `func ToUpper(s string) string` | Converts to uppercase. |
| **strconv** | `Atoi` | `func Atoi(s string) (int, error)` | String to Integer. |
| **strconv** | `Itoa` | `func Itoa(i int) string` | Integer to String. |
| **strconv** | `ParseBool` | `func ParseBool(s string) (bool, error)` | String to Boolean. |
| **strconv** | `ParseFloat`| `func ParseFloat(s string, bit int) (float64, error)` | String to Float. |

---

## 1. What is a String in Go?

In Go, a string is an **immutable** sequence of bytes. This means that once a string is created, its contents cannot be changed.

```go
s := "hello, world"
// s[0] = 'H' // This will cause a compile-time error
```

Strings are typically UTF-8 encoded, but this is not enforced by the language itself. A string can contain arbitrary bytes.

A string is essentially a read-only slice of bytes. You can get its length using the `len()` function, which returns the number of bytes (not characters or runes).

```go
s := "hello"
fmt.Println(len(s)) // 5

s2 := "你好"
fmt.Println(len(s2)) // 6 (each Chinese character is 3 bytes in UTF-8)
```

## 2. Runes: The Go Way of Handling Characters

To handle individual characters, especially in a Unicode-aware way, Go uses the `rune` type. A `rune` is an alias for `int32` and represents a Unicode code point.

When you iterate over a string using a `for...range` loop, you get runes, not bytes.

```go
s := "你好"
for index, r := range s {
    fmt.Printf("index: %d, rune: %c, value: %v\n", index, r, r)
}
// Output:
// index: 0, rune: 你, value: 20320
// index: 3, rune: 好, value: 22909
```

Notice that the index jumps from 0 to 3. This is because the `for...range` loop decodes one UTF-8 encoded rune at a time, and the index is the starting byte position of that rune.

To work with strings as a sequence of runes, you can convert a string to a slice of runes:

```go
runes := []rune("你好")
fmt.Println(len(runes)) // 2
```

## 3. The `strings` Package

The `strings` package is your go-to tool for string manipulation. It provides a wide array of functions, including:

-   **Searching:**
    -   `strings.Contains(s, substr)`: Checks if a substring is present.
    -   `strings.HasPrefix(s, prefix)`: Checks for a prefix.
    -   `strings.HasSuffix(s, suffix)`: Checks for a suffix.
    -   `strings.Index(s, substr)`: Finds the first index of a substring.

-   **Modification (returns a new string):**
    -   `strings.ToUpper(s)`: Converts to uppercase.
    -   `strings.ToLower(s)`: Converts to lowercase.
    -   `strings.TrimSpace(s)`: Removes leading/trailing whitespace.
    -   `strings.Replace(s, old, new, n)`: Replaces occurrences of a substring.
    -   `strings.ReplaceAll(s, old, new)`: Replaces all occurrences.

-   **Joining and Splitting:**
    -   `strings.Join(elems []string, sep string)`: Joins a slice of strings.
    -   `strings.Split(s, sep string)`: Splits a string by a separator.

-   **`strings.Builder`:**
    For building strings piece by piece, using `strings.Builder` is much more efficient than repeated string concatenation (`+`).

    ```go
    var b strings.Builder
    for i := 0; i < 10; i++ {
        b.WriteString(fmt.Sprintf("%d ", i))
    }
    fmt.Println(b.String())
    ```

## 4. The `strconv` Package

The `strconv` package (string conversion) is used for converting strings to and from basic data types.

-   `strconv.Itoa(i int)`: Converts an integer to a string.
-   `strconv.Atoi(s string)`: Converts a string to an integer.
-   `strconv.ParseFloat(s string, bitSize int)`: Converts a string to a float.
-a  `strconv.FormatFloat(f float64, fmt byte, prec, bitSize int)`: Converts a float to a string.
-   `strconv.ParseBool(s string)`: Converts a string to a boolean.
-   `strconv.FormatBool(b bool)`: Converts a boolean to a string.

## 5. Reversing a String (The Unicode Challenge)

Reversing a string is a classic problem that highlights the complexities of working with Unicode. A naive byte-wise reversal will corrupt multi-byte characters.

The `ReverseString.go` file likely contains a correct implementation that handles Unicode properly. The correct way to reverse a string in Go is to convert it to a slice of runes, reverse the slice, and then convert it back to a string.

```go
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    fmt.Println(Reverse("hello")) // "olleh"
    fmt.Println(Reverse("你好"))   // "好你"
}
```

This approach ensures that each Unicode code point remains intact, and the reversal happens at the character level, not the byte level.

## Conclusion

Working with strings in Go is generally straightforward, but it's crucial to understand the distinction between bytes and runes to handle international text correctly. The `strings` package provides most of the tools you'll need for common string manipulations, while `strconv` handles conversions between strings and other basic types. Always be mindful of Unicode when your application needs to process text from around the world.

# Stacks in Go

This guide explores the stack data structure in Go.

## Quick Reference: Stack Operations (Slice-based)

| Operation | Implementation (Generic) | Description | Time Complexity |
| :--- | :--- | :--- | :--- |
| **Push** | `s = append(s, v)` | Adds `v` to the top of the stack. | O(1) amortized |
| **Pop** | `v, s = s[len(s)-1], s[:len(s)-1]` | Removes and returns the top element. | O(1) |
| **Peek** | `v = s[len(s)-1]` | Returns top element without removal. | O(1) |
| **IsEmpty** | `len(s) == 0` | Returns true if stack has no elements. | O(1) |
| **Size** | `len(s)` | Returns number of elements in stack. | O(1) |

---

## 1. What is a Stack?

A stack is a linear data structure that follows the **LIFO (Last-In, First-Out)** principle. This means that the last element added to the stack will be the first one to be removed.

Think of a stack of plates: you add a new plate to the top, and when you need a plate, you take it from the top.

The main operations on a stack are:

-   **Push:** Add an element to the top of the stack.
-   **Pop:** Remove the element from the top of the stack.
-   **Peek (or Top):** Look at the top element of the stack without removing it.
-   **IsEmpty:** Check if the stack is empty.

## 2. Implementing a Stack in Go

Go doesn't have a built-in stack type in its standard library. However, it's straightforward to implement one using a slice. A slice provides all the necessary functionality to create an efficient stack.

The `arrayStack.go` file provides an example of a stack implemented with a slice.

Here's how you can define a simple stack for integers:

```go
package main

import "fmt"

// Stack represents a stack that holds a slice of integers.
type Stack struct {
    items []int
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

// Pop removes and returns the top item of the stack.
// It returns an error if the stack is empty.
func (s *Stack) Pop() (int, error) {
    if len(s.items) == 0 {
        return 0, fmt.Errorf("stack is empty")
    }
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, nil
}

// Peek returns the top item of the stack without removing it.
func (s *Stack) Peek() (int, error) {
    if len(s.items) == 0 {
        return 0, fmt.Errorf("stack is empty")
    }
    return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}
```

### Using Generics for a Reusable Stack (Go 1.18+)

With the introduction of generics in Go 1.18, we can create a reusable stack that can hold any type of data.

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
    if len(s.items) == 0 {
        var zero T
        return zero, fmt.Errorf("stack is empty")
    }
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, nil
}
// ... and so on for Peek, IsEmpty
```

This generic implementation is much more flexible and type-safe.

## 3. Alternative: Using `container/list`

While using a slice is the most common and often the most efficient way to implement a stack in Go, you can also use the `container/list` package (a doubly linked list).

```go
import "container/list"

// A stack implemented using a linked list.
type Stack struct {
    list *list.List
}

func (s *Stack) Push(value interface{}) {
    s.list.PushBack(value)
}

func (s *Stack) Pop() (interface{}, error) {
    if s.list.Len() == 0 {
        return nil, fmt.Errorf("stack is empty")
    }
    e := s.list.Back()
    s.list.Remove(e)
    return e.Value, nil
}
```

**Slice vs. Linked List for a Stack:**

-   **Slice-based Stack:**
    -   **Pros:** More memory efficient as it stores elements in a contiguous block of memory, which is better for CPU caching (locality of reference). Generally faster for most use cases.
    -   **Cons:** Appending may require re-allocation and copying of the underlying array if the capacity is exceeded. However, Go's `append` strategy amortizes this cost effectively.

-   **Linked List-based Stack:**
    -   **Pros:** Push and Pop operations are guaranteed O(1) time complexity without the need for re-allocation.
    -   **Cons:** Less memory efficient due to the overhead of pointers in each node. Can be slower in practice due to poor cache performance.

For these reasons, **a slice-based implementation is the idiomatic and recommended approach for building a stack in Go.**

## 4. Use Cases for Stacks

Stacks are used in a wide variety of algorithms and programming scenarios:

-   **Function Calls:** The call stack manages active function calls in a program.
-   **Undo/Redo Functionality:** Stacks can be used to store a history of actions.
-   **Expression Evaluation:** Converting infix expressions to postfix (Reverse Polish Notation) and evaluating them.
-   **Backtracking Algorithms:** Such as maze-solving or in parsing.
-   **Depth-First Search (DFS):** An iterative implementation of DFS uses a stack to keep track of nodes to visit.

## Conclusion

While Go doesn't provide a built-in stack type, implementing one using a slice is simple, efficient, and idiomatic. By leveraging Go's generics, you can create a type-safe, reusable stack that can be a valuable tool in your programming toolkit.

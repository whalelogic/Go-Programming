# Linked Lists in Go

This guide explores the concept of linked lists in Go, covering both custom implementations and the use of the standard library's `container/list` package.

## 1. What is a Linked List?

A linked list is a linear data structure where elements are not stored at contiguous memory locations. Instead, each element (a *node*) contains a value and a pointer to the next node in the sequence.

There are two main types of linked lists:

-   **Singly Linked List:** Each node points only to the next node.
-   **Doubly Linked List:** Each node points to both the next and the previous node.

Linked lists offer several advantages over arrays (and slices in Go):

-   **Dynamic Size:** They can grow and shrink easily.
-   **Efficient Insertions/Deletions:** Inserting or deleting an element in the middle of the list is fast (O(1)) if you have a pointer to the node.

However, they also have disadvantages:

-   **No Random Access:** Accessing an element by index requires traversing the list from the beginning (O(n)).
-   **Extra Memory:** Each node requires extra memory to store the pointer(s).

## 2. Custom Linked List Implementation

The `linkedList.go` file provides an example of a custom singly linked list implementation in Go.

A typical node structure for a singly linked list looks like this:

```go
type Node struct {
    Value int
    Next  *Node
}
```

And the list itself would be represented by a pointer to the head (the first node) of the list:

```go
type LinkedList struct {
    Head *Node
}
```

Common operations on a linked list include:

-   **Insert:** Adding a new node to the list (at the beginning, end, or in the middle).
-   **Delete:** Removing a node from the list.
-   **Search:** Finding a node with a specific value.
-   **Traverse:** Iterating through all the nodes in the list.

While implementing a linked list from scratch is a great way to understand the data structure, for most practical purposes, you should use the standard library's implementation.

## 3. The `container/list` Package

Go's standard library provides a doubly linked list implementation in the `container/list` package. This implementation is robust, well-tested, and ready for production use.

### Creating a New List

You can create a new list using `list.New()`:

```go
import "container/list"

l := list.New()
```

### Adding Elements

The `container/list` package provides several functions for adding elements:

-   `PushFront(v interface{}) *Element`: Adds an element to the front of the list.
-   `PushBack(v interface{}) *Element`: Adds an element to the back of the list.
-   `InsertBefore(v interface{}, mark *Element) *Element`: Inserts an element before another element.
-   `InsertAfter(v interface{}, mark *Element) *Element`: Inserts an element after another element.

```go
l := list.New()
l.PushBack("world")
l.PushFront("hello") // List is now: "hello", "world"
```

### Traversing a List

You can traverse a list using a `for` loop, starting from the `Front()` or `Back()` element and following the `Next()` or `Prev()` pointers.

```go
for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
}
```

### Removing Elements

The `Remove(e *Element) interface{}` function removes an element from the list.

```go
e := l.Front()
l.Remove(e) // Removes "hello"
```

## 4. When to Use a Linked List in Go

In Go, slices are often a better choice than linked lists due to their cache-friendliness and ease of use. Slices are dynamic arrays that can grow and shrink, and they provide O(1) random access.

However, linked lists can be a good choice in specific scenarios:

-   **Frequent Insertions/Deletions in the Middle:** If your application involves frequent insertions or deletions in the middle of a large collection, a linked list can outperform a slice.
-   **Implementing Other Data Structures:** Linked lists are often used as building blocks for other data structures, such as queues, stacks, and hash maps.
-   **Need for Stable Pointers:** When you need pointers to elements that remain valid even after other elements are inserted or deleted, a linked list can be useful.

## Conclusion

Understanding linked lists is essential for any software engineer. While you may not need to implement one from scratch in your day-to-day Go programming, knowing when and how to use the `container/list` package can be a valuable tool in your arsenal. Always consider the trade-offs between linked lists and slices to choose the most appropriate data structure for your specific use case.

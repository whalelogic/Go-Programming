# 🧠 Go Programming

This repository demonstrates core **data structures** and **algorithms** implemented in **Go (Golang)** — with a focus on performance, clarity, and idiomatic code.

---

## 📁 Repository Structure

```
Go-Programming/
├── main.go                        # Entry point (stack, linked list, sorting, recursion demos)
├── go.mod / go.sum                # Module definition and dependency checksums
│
├── docs/                          # All documentation, organized by topic
│   ├── data-structures/           # Linked list, maps, slices, stacks, tree
│   ├── algorithms/                # Sorting
│   ├── concurrency/               # Channels, goroutines, mutexes
│   ├── networking/                # HTTP, TCP reference
│   ├── strings/                   # String manipulation
│   ├── utilities/                 # Utility function reference
│   └── stdlib/                    # Go standard library reference (errors, I/O, OS, etc.)
│
├── lists/                         # Generic linked list package
├── sort/                          # Bubble sort package
├── stack/                         # Stack (linked-list-backed) package
├── recursion/                     # Factorial and Fibonacci functions
├── tree/                          # Binary search tree / TreeSort package
├── utilities/                     # Shared utility functions (max, swap, reverse, sum, etc.)
├── web/                           # HTTP server package (Gorilla Mux)
├── network/json/                  # HTTP JSON fetch package
│
├── examples/                      # Standalone runnable programs
│   ├── concurrency/               # Goroutines, channels (simple & advanced), select
│   ├── data-structures/           # Slice operations, map operations
│   ├── algorithms/                # Filter, recursion demos
│   ├── networking/                # URL fetcher
│   ├── fileops/                   # File system operations
│   ├── strings/                   # String manipulation demo
│   ├── user-input/                # stdin user input demo
│   └── memory/                    # new() vs make() allocation demo
│
└── data/                          # Sample data files (JSON, CSV, TXT)
```

---

## 📚 Packages

| Package | Path | Description |
|---------|------|-------------|
| `lists` | `lists/` | Generic doubly-linked list with push, pop, insert, iterate |
| `sort` | `sort/` | Bubble sort implementation |
| `stack` | `stack/` | Stack backed by linked list nodes |
| `recursion` | `recursion/` | Factorial and Fibonacci (recursive & iterative) |
| `tree` | `tree/` | Binary search tree with in-order traversal / TreeSort |
| `utilities` | `utilities/` | Max, swap, reverse, sum, filter, string reader, URL parser |
| `web` | `web/` | HTTP server with graceful shutdown (Gorilla Mux) |
| `json` | `network/json/` | HTTP JSON fetch helpers |

---

## 🚀 Example: TreeSort

```go
values := []int{5, 1, 9, 2, 6}
TreeSort(values)
fmt.Println("Sorted:", values)
```
<img width="909" height="895" alt="image" src="https://github.com/user-attachments/assets/dc4fd440-9f6b-482a-ae82-0ff0397a60be" />




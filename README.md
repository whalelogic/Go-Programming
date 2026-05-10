# 🧠 Go Programming: Learning & Reference

This repository is a comprehensive learning resource for **Go (Golang)**, featuring core language concepts, data structures, algorithms, and advanced topics. It focuses on performance, clarity, and idiomatic code, with detailed technical reference tables for the standard library.

---

## 📚 Repository Structure

The repository is organized by topic to make finding specific information and examples easy.

### 1. [Concepts](./concepts/)
Core language fundamentals and mechanics.
- **[Variables](./concepts/VARIABLES.md)**: Declarations, types, and scope.
- **[Functions](./concepts/FUNCTIONS.md)**: Technical table of function features, closures, and variadic parameters.
- **[Structs](./concepts/STRUCTS.md)**: Custom types and composition.
- **[Interfaces](./concepts/INTERFACES.md)**: Duck typing and decoupling in Go.
- **[Errors](./concepts/ERRORS.md)**: Idiomatic error handling patterns.
- **[Language Spec](./concepts/spec/)**: Official Go specification files.

### 2. [Data Structures](./data_structures/)
Built-in and custom data structure implementations.
- **[Slices](./data_structures/Slices.md)**: Reference table for slice operations and the `slices` package.
- **[Maps](./data_structures/Maps.md)**: Reference table for map operations and the `maps` package.
- **[Strings](./data_structures/Strings.md)**: Reference table for `strings` and `strconv` packages.
- **[Linked Lists](./data_structures/Linked_List.md)**: Custom implementations and `container/list`.
- **[Stacks](./data_structures/Stacks.md)**: Slice-based and linked-list stack implementations.
- **[Trees](./data_structures/Tree.md)**: Binary trees, BSTs, and traversal algorithms.

### 3. [Algorithms](./algorithms/)
Common algorithms implemented in Go.
- **[Sorting](./algorithms/sorting/)**: Quick Sort, Merge Sort, and the `sort` package reference.
- **[Searching](./algorithms/searching/)**: Binary search and other search techniques.
- **[Recursion](./algorithms/recursion/)**: Factorials, Fibonacci, and recursive patterns.

### 4. [Standard Library](./standard_library/)
Detailed guides and technical tables for Go's standard packages.
- **[Built-in Functions](./standard_library/builtins/BUILTINS.md)**: **Full technical table** of all Go built-ins (`append`, `make`, `len`, `clear`, etc.).
- **[OS](./standard_library/OS.md)**: File system, environment variables, and process management.
- **[IO](./standard_library/IO.md)**: Readers, writers, and stream manipulation.
- **[Utilities](./standard_library/utilities/)**: Common helper functions, filters, and file operations.

### 5. [Advanced Topics](./advanced/)
Complex systems and concurrent programming.
- **[Concurrency](./advanced/concurrency/)**: Goroutines, channels, and synchronization primitives.
- **[Networking](./advanced/networking/)**: TCP/UDP, JSON fetching, and URL parsing.
- **[Web](./advanced/web/)**: HTTP servers, templates, and REST API examples.

---

## 🚀 Getting Started

Each directory contains a `main.go` or specific example files. You can run them directly using the Go tool:

```bash
# Example: Run the built-in functions demo
go run standard_library/builtins/main.go
```

## 🛠 Features
- **Technical Tables**: Most documentation files include "Quick Reference" tables for functions, methods, and types.
- **Idiomatic Code**: Examples follow Go best practices and utilize modern features (Go 1.21+ where applicable).
- **Comprehensive Coverage**: From basic variables to advanced concurrency and networking.

---
*Happy Coding in Go!*

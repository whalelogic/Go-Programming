// Package tree implements a binary search tree and provides a function to sort a slice of integers using tree sort algorithm.
package tree


type Tree struct {
	val         int
	left, right *Tree
}


func TreeSort(values []int) []int {
	var root *Tree
	for _, v := range values {
		root = root.Add(v)
	}
	values = appendValues(values[:0], root)
	return values
}

func (t *Tree) Add(val int) *Tree {
	if t == nil {
		return &Tree{val: val}
	}
	if val < t.val {
		t.left = t.left.Add(val)
	} else {
		t.right = t.right.Add(val) 
	}
	return t
}

// function
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.val)
		values = appendValues(values, t.right)
	}
	return values
}

// Go's hybrid OOP/Functional model suggests using methods over pure functions in this instance. 

func (t *Tree) AppendValues(val []int) []int {
	if t != nil {
		val = t.left.AppendValues(val)
		val = append(val, t.val)
		val = t.right.AppendValues(val) 
	}
	return val
}



// Example main.go
//
// func main() {
// 	arr := make([]int, 10)
//
// 	for i := range arr {
// 		arr[i] = rand.Intn(100)
// 	}
//
// 	for _, val := range arr {
// 		fmt.Printf("Value: %v\n", val)
// 		r := reflect.TypeOf(val)
// 		fmt.Printf("\n\rType: %v\n", r)
// 	}
//
// 	TreeSort(arr)
// 	fmt.Println("Sorted:", arr)
// }

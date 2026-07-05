package main


func DivisibleByTwo(n []int) []int {
	i := 0	
	for _, v := range n {
		if v%2 == 0 {
			n[i] = v 
			i++
		}
	}
	return n[:i]
}

//
// func main() {
//
// 	masterList := []int{4,3,2,1,3,4,4,3,3,32,44}
// 	// lst is a slice with a length of 9.
// 	lst := []int{33,4,5,2,12,55,6,10,44}
//
//
// 	for i, v := range lst {
// 		fmt.Printf("Index: %v Value: %v\n", i, v)
// 	}
//
// 	slices.Sort(lst)
// 	fmt.Println("Sorted: ", lst)
//
//
//   f := DivisibleByTwo(masterList)
// 	fmt.Println("Filtering masterList (divisible by 2): ", f)
// }

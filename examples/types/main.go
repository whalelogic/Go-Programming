package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num int = 44
	var str string = "Keith"
	var floatNum float64 = 3.14
	var boolVal bool = true
	var complexNum complex128 = 1 + 2i
	var arr [3]int = [3]int{1, 2, 3}
	var slice []int = []int{4, 5, 6}
	var m map[string]int = map[string]int{"one": 1, "two": 2}
	var ch rune = 'A'
	var fn func(int) int = func(x int) int { return x * x }
	var ptr *int = &num
	var iface interface{} = "Hello, World!"
	var nilVal interface{} = nil
	var i any = 42

	// print the types of the variables
	fmt.Println("Types: ", reflect.TypeOf(num), reflect.TypeOf(str), reflect.TypeOf(floatNum), reflect.TypeOf(boolVal), reflect.TypeOf(complexNum), reflect.TypeOf(arr), reflect.TypeOf(slice), reflect.TypeOf(m), reflect.TypeOf(ch), reflect.TypeOf(fn), reflect.TypeOf(ptr), reflect.TypeOf(iface), reflect.TypeOf(nilVal))

	// values 
	fmt.Println("Values: ", num, str, floatNum, boolVal, complexNum, arr, slice, m, ch, fn(5), *ptr, iface, nilVal)

	s, ok := i.(string) 
	fmt.Println("Value of I:", i)
	fmt.Println("Type of S: ", reflect.TypeOf(s))
	fmt.Println("Type of I: ", reflect.TypeOf(i))
	if !ok {
		fmt.Printf("i: %v is not a string!\n", i)
	}
	si := i.(int)
	fmt.Println("Value of SI:", rune(si))
	

}

package main

import (
	"fmt"
	"reflect"

	"github.com/whalelogic/Go-Programming/stack"
	"github.com/whalelogic/Go-Programming/utils"
)



func main() {

	arr := stack.Stack{}
	arr.Push(10)
	arr.Push(20)
	arr.Push(30)
	arr.Pop()
	arr.Print()
	arr.Peek()

	fmt.Println(reflect.TypeOf(arr))

	slc := []int{5, 8, 13, 21, 34, 55, 89, 144}
	slc2 := []int{13,21,34,55}
	fmt.Println(slc)
	fmt.Println(utils.GenerateFibs(slc))
	fmt.Println(utils.GenerateFibs(slc2))


	

}

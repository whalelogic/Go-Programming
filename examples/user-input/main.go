package main

import (
	"fmt"
) 


type Person struct {
	Name string
	Age  int
}

func GetPerson() Person {
	var p Person
	fmt.Print("Enter name: ")
	fmt.Scan(&p.Name)
	fmt.Print("Enter age: ")
	fmt.Scan(&p.Age)
	return p
}

func main() {

	p := GetPerson()
	fmt.Printf("Hello, %s! You are %d years old.\n", p.Name, p.Age)


}

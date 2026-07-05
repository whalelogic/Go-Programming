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


func (p Person) Greet() {
	fmt.Printf("Hello, %s! You are %d years old.\n", p.Name, p.Age)
}


func GetUserInput() ([]string, error) {
	var inputs []string
	var input string

	fmt.Println("Enter strings (type 'exit' to finish):")
	for {
		fmt.Print("> ")
		fmt.Scan(&input)
		if input == "exit" {
			break
		}
		inputs = append(inputs, input)
	}
	return inputs, nil
}




func main() {

	p := GetPerson()
	p.Greet()

	inputs, err := GetUserInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	
	fmt.Println("You entered:")
	for _, input := range inputs {
		fmt.Println(input)
	}



}

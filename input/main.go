package main

import (
	"fmt"
	"io"
	"os"
) 


func GetUserInput(s any) any {
	r, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))
	return string(r)
}

func main() {
	g := GetUserInput("Type your name: ")
	print(g)
}

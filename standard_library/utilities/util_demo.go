package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/whalelogic/Go-Programming/standard_library/utilities/fileops"
)


func main() {
	
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer f.Close()
	var b bytes.Buffer
	b.ReadFrom(f)
	b.WriteTo(os.Stdout)
	b.WriteString("\nThis is a new line added to the buffer.\n")
	bn := b.Available()
	fmt.Println("Available bytes in buffer:", bn)
	b.WriteByte('Z')
	bn = b.Available()
	fmt.Println("Available bytes in buffer after writing a byte:", bn)
	
	bytesToString := b.String()
	fmt.Println("File content:", bytesToString)

	_ = fileops.WriteFile("output.txt", bytesToString)
}

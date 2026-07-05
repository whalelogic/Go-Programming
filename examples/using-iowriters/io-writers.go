package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}



func main() {

	var w io.Writer = os.Stdout
 	fmt.Fprintf(w, "Hello, World!\n")

	var r io.Reader = os.Stdin
	buf := make([]byte, 1024)

	n, err := r.Read(buf)
	check(err)
	fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))

	l, le := cap(buf), len(buf[:n])
	fmt.Println("Length of input:", l, "Length of buffer:", le)

	for scanner := buf[:n]; len(scanner) > 0; {
		fmt.Printf("Scanner: %s\n", string(scanner))
		scanner = scanner[:len(scanner)-1]
	}

}

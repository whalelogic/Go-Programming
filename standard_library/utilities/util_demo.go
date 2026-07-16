package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/whalelogic/Go-Programming/standard_library/utilities/fileops"
)

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}


func hasMDext(path string) (bool, error) {
	ext := filepath.Ext(path)
	if ext == ".md" {
		fmt.Println("Definitely Markdown.")
		return true, nil
	} else {
		fmt.Println("Not a markdown file.")
		return false, nil
	}
}

//
// func isMarkdown(bs *bufio.Scanner) bool {
// 	m := false
// 	for bs.Scan() {
// 		line := bs.Text()
// 		if len(line) > 0 {
// 			fb := line[0]
// 			if fb == '#' {
// 				m = true
// 				fmt.Println("Contains markdown.")
// 			}
// 		}
// 	}
// 	return m
// }


func main() {

	path := "./test.txt"
	mdPath := "./Utility_Functions.md"
	h, e := hasMDext(mdPath)
	check(e)
	fmt.Println("has .md extension? ", h)
	a := fileops.IsMarkdown(mdPath)
	fmt.Println("a: ", a)

	h, e = hasMDext(path)
	check(e)
	fmt.Println("has .md extension? ", h)

	// f is a string containing all of the files data 
	f, err := fileops.ReadFile(path)
	check(err)

	// fstr is a *strings.Reader
	// it has several helpful methods like ReadAt, ReadByte, len, peek, size.
	fstr := strings.NewReader(f)
	fmt.Println("Length: ", fstr.Len())
	fmt.Println("Size: ", fstr.Size())

	scanner := bufio.NewScanner(fstr)
	for scanner.Scan() {
		fmt.Println("lines: ", scanner.Text())
	}




}

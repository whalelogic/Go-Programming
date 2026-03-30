package filops


import (
	"io"
	"testing/fstest"
	"io/fs"
	"time"
)


func DoSomethingWithFile() {

	fsys := fstest.MapFS{
		"test.txt": 	{},
		"readme.md": 	{},
		"main.go":  	{},
	}
	var patterns = []string{"*.txt", "*.md", "*.go"}

	file, err := fsys.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	for _, p := range(patterns) {
		matched, err := fs.Glob(fsys, p)
		if err != nil {
			panic(err)
		}
		println("Matched files for pattern", p, ":")
		for _, m := range matched {
			println(" -", m)
		}
	}

	// Print file information
	println("Name:", info.Name())
	println("Size:", info.Size())
	println("Mode:", info.Mode().String())
	println("ModTime:", info.ModTime().Format(time.RFC3339))
	println("IsDir:", info.IsDir())

	// Read file content
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	println("Content:", string(content))

}

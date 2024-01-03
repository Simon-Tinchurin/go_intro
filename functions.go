package main

import (
	"fmt"
	"os"
)

// best practice - capitalize name of a function, except those that in package main

// function name( param and type) returns (result or error)
// always returns 2 parameters
func getFiles(path string) ([]string, error) {
	// define variable with type
	var files []string
	// go to the path with ReadDir
	dirFiles, err := os.ReadDir(fmt.Sprintf("%s", path))
	// nil is like None in python
	if err != nil {
		return nil, err
	}
	// go through dirFiles
	for _, i := range dirFiles {
		// and add file name to slice
		files = append(files, i.Name())
	}
	return files, nil
}

// -----------------
func test(path string) ([]string, error) {
	return getFiles(path)
}

// ----------------
func returnRog() (rog int) {
	rog = 123
	return
}

func functions_main() {
	// write function in a variable
	functionInVariable := func(x, y int) int {
		i := (x + y) * 2
		return i
	}

	f, _ := getFiles(".")
	fmt.Println(f)
	fmt.Println(returnRog())
	fmt.Println(functionInVariable(1, 5))
}

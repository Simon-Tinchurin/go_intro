package main

import "fmt"

// Number of the lesson: 4

// declearing variables
var (
	// declare a variable
	name = "name"
	// declare a variable with type
	firstname string = "Foo"
	// declare an empty variable
	lastName string
)

// declearing constants, same method for a few constants like in vars
// always in lowercase
const someConstant = 150

func main() {
	var version int
	version = 100            // infer integer
	otherVersion := "second" // infer string
	anotherVersion := 10.1   // infer float
	// stands for format print line
	fmt.Println(name, version, otherVersion, anotherVersion, someConstant)
}

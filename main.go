package main

import "fmt"

var (
	// declare a variable
	name = "name"
	// declare a variable with type
	firstname string = "Foo"
	// declare an empty variable
	lastName string
)

func main() {
	version := 1             // infer integer
	otherVersion := "second" // infer string
	anotherVersion := 10.1   // infer float

	fmt.Println(version, otherVersion, anotherVersion)
}

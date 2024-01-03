package main

import "fmt"

// Control structures
// for Loops
// range
// break
func les_6_main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	names := []string{"a", "b", "c", "d"}

	// iterate with for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	//iterate with range
	for i, name := range names {
		if name == "c" {
			break
		}
		fmt.Println(i, name)
	}
	fmt.Println("Break out of the loop")
}

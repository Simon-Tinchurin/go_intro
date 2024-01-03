package main

import (
	"fmt"
	"sort"
)

func arrauys_main() {
	// empty array of strings with len 3
	var array [3]string

	newArray := [4]int{1, 2, 3, 4}

	strArray := []string{"a", "d", "c", "f", "a"}

	// dunamic array
	anotherArray := []int{30, 16, 82, 12}
	// add to dynamic array
	anotherArray = append(anotherArray, 500)
	// another way to create an array
	againArray := make([]int, 10)

	//sort array
	sort.Ints(anotherArray)
	sort.Strings(strArray)

	fmt.Print(anotherArray, "\n")

	fmt.Print(againArray, "\n")

	fmt.Print(strArray, "\n")

	// iterate through array with 'for' loop
	for _, val := range newArray {
		fmt.Println(val)
	}

	// get length of an array like in python - len()
	fmt.Println(len(array))
	fmt.Println(array)
}

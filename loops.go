package main

import (
	"fmt"
	"log"
	"os"
)

// in go is only 'for' loop
// there's no 'while' like in python

// i++ is the same as i = i + 1

func loops_main() {
	i := 1
	for i <= 2 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j <= 2; j++ {
		fmt.Println(j)
	}

	text := []string{"one", "two", "three"}
	// in iteration through list or smth first
	// parameter is always an index, so we use '_'
	for _, value := range text {
		if value == "two" {
			// there's also things like 'break' and 'continue'
			continue
		}
		fmt.Println(value)
	}

	fmt.Println("#######")

	files, err := os.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	// if we need infinite loop (there is no 'while' in go)
	// like 'while=True' in python
	// for {
	// 	fmt.Println("asd")
	// }

}

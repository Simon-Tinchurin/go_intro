package main

import (
	"fmt"
	"time"
)

// goroutines

func lesson_14() {
	// result := fetchResource(1)
	// fmt.Println("the result: ", result)

	// append 'go' to the front of the function to create a goroutine from this:
	go fetchResource(1) // async call
	// schedule an anonymous function in a separate goroutine
	go func() {
		fetchResource(2)
	}()
	// same methods above

	// channel in golang will always blocks if its full,
	// there are 2 types - unbuffered and buffered.
	// declare channel (append "ch" at the end of a variable - common practice):
	// resultch := make(chan string) // -> unbuffered channel

	resultch := make(chan string) // -> buffered channel, 10 is the size of a buffer
	// reading from the channel into a variable:
	go func() {
		result := <-resultch
		fmt.Println(result)
	}()
	// write things into that channel:
	resultch <- "foo" // -> is now FULL -> IT WILL BLOCK -> code below will never execute

}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}

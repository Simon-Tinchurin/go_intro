package main

import "fmt"

// channels

func lesson_15() {
	// declaring channel:
	msgch := make(chan string, 128)
	// writing to a channel:
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	// close the channel to avoid the error:
	// fatal error: all goroutines are asleep - deadlock!
	close(msgch)

	// ranging over a channel (this is our consumer):
	// for msg := range msgch {
	// 	fmt.Println("the message is:", msg)
	// }
	// fmt.Println("Done reading all the messages from the channel!")

	for { // while
		// check if channel is closed to break the loop ("ok" is bool)
		msg, ok := <-msgch
		if !ok {
			break
		}
		fmt.Println("the message is:", msg)
	}
}

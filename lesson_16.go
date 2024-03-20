package main

import (
	"fmt"
	"time"
)

// control flow and syncronization examples

type Server2 struct {
	quitch chan struct{} // binary
	msgch  chan string
}

// to initialize new server
func newServer() *Server2 {
	return &Server2{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server2) start() {
	fmt.Println("server starting")
	s.loop() // block
}

func (s *Server2) handleMessage(msg string) {
	fmt.Println("we recieved a message:", msg)
}

func (s *Server2) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server2) quit() {
	// close(s.quitch)
	s.quitch <- struct{}{}
}

func (s *Server2) loop() {
mainLoop:
	for {
		select {
		case <-s.quitch:
			// do some stuff when we need to quit
			fmt.Println("server is shutting down")
			break mainLoop
		case msg := <-s.msgch:
			// do some stuff when we have message
			s.handleMessage(msg)
			// default:
		}
	}
}

func lesson_16() {
	s := newServer()
	// go s.start()

	// for i := 0; i < 10; i++ {
	// 	s.sendMessage(fmt.Sprintf("handle this number - %d", i))
	// }
	// time.Sleep(time.Second * 5)

	go func() {
		time.Sleep(time.Second * 5)
		s.quit()
	}()
	s.start()
}

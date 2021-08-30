package main

import "fmt"

type Server struct {
	id   int
	name string
}

// This is also called a method
func (server Server) display() {
	fmt.Printf("Server Id: %v\nServer Name: %v\n", server.id, server.name)
}

func main() {
	s := Server{
		id:   1,
		name: "K.R.I.S.P.Y.",
	}
	s.display()
}

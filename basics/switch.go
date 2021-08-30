package main

import (
	"fmt"
	"log"
)

func main() {
	var (
		a int
	)
	if _, err := fmt.Scanf("%v", &a); err != nil {
		log.Panic(err)
	}
	switch a {
	case 1, 2, 3:
		fmt.Println("One, Two or Three")
	case 4, 5, 6:
		fmt.Println("Four, Five, Six")
	default:
		fmt.Println("Another Number")
	}
}

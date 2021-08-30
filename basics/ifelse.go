package main

import (
	"fmt"
	"log"
)

func main() {
	if true {
		fmt.Println("This is True")
	}
	first := map[string]int{
		"hishaam":  19,
		"hasan":    15,
		"sharanya": 18,
		"mayank":   19,
	}
	if pop, ok := first["mayank"]; ok {
		fmt.Println(pop)
	}
	var (
		a int
		b int
	)
	if _, err := fmt.Scanf("%v", &a); err != nil {
		log.Panic("Error")
	}
	if _, err := fmt.Scanf("%v", &b); err != nil {
		log.Panic("Error")
	}
	if a == b {
		fmt.Println("Equal")
	} else if a > b {
		fmt.Println("A is greater")
	} else if a < b {
		fmt.Println("B is greater")
	}
}

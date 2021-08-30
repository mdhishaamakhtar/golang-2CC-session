package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%v", &n); err != nil {
		log.Panic(err)
	}
	// basic for loops
	fmt.Println("Printing First For Loop")
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
	// while loops is also for loops like this
	fmt.Println("Printing Second For Loop")
	for n < 20 {
		fmt.Println(n)
		n++
	}

	list := []int{10, 20, 30}
	fmt.Println("Printing Range based loop for array")
	for k, v := range list {
		fmt.Println(k, v)
	}

	mp := map[string]int{
		"hishaam":  10,
		"hasan":    20,
		"sharanya": 30,
	}
	fmt.Println("Printing Range based loop for map")
	for k, v := range mp {
		fmt.Println(k, v)
	}

	str := "Hishaam"
	fmt.Println("Printing Range based loop for string")
	for k, v := range str {
		fmt.Println(k, v)
	}

	fmt.Println("Printing Infinite for loop with a break")
	for {
		if n > 35 {
			break
		} else {
			fmt.Println(n)
			n++
		}
	}
}

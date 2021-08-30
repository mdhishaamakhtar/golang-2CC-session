package main

import (
	"fmt"
	"log"
)

func main() {
	first := complex(10, 20)
	second := "GoLang is amazing"
	third := 3.141
	fourth := true
	var fifth int64
	fifth = 10
	var (
		sixth   = 10
		seventh = "hello"
	)
	var eighth int
	_, err := fmt.Scanf("%v", &eighth)
	if err != nil {
		log.Panic(err)
	}
	ninth := fmt.Sprintf("My name is %v", "hishaam")
	fmt.Printf("%v, %T\n", first, first)
	fmt.Printf("%v, %T\n", second, second)
	fmt.Printf("%v, %T\n", third, third)
	fmt.Printf("%v, %T\n", fourth, fourth)
	fmt.Printf("%v, %T\n", fifth, fifth)
	fmt.Printf("%v, %T\n", sixth, sixth)
	fmt.Printf("%v, %T\n", seventh, seventh)
	fmt.Printf("%v, %T\n", eighth, eighth)
	fmt.Printf("%v, %T\n", ninth, ninth)
}

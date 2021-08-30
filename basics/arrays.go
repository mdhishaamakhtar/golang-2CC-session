package main

import (
	"fmt"
)

func main() {
	var first = [4]int{0, 10, 20, 30}
	second := [4]int{0, 10, 20, 30}
	third := [...]int{0, 10, 20, 30, 40, 50, 60}
	fourth := []int{0, 10, 20, 30, 40, 50, 60, 70}
	fifth := third[3:6]
	fourth = append(fourth, 80, 90, 100)
	fourth = append(fourth[:3], fourth[4:]...)
	sixth := fourth[2:7]
	seventh := make([]int, 3)
	seventh = []int{0, 10, 20}
	fmt.Printf("%v\t%T\tlength: %v\n", first, first, len(first))
	fmt.Printf("%v\t%T\tlength: %v\n", second, second, len(second))
	fmt.Printf("%v\t%T\tlength: %v\n", third, third, len(third))
	fmt.Printf("%v\t%T\tlength: %v\n", fourth, fourth, len(fourth))
	fmt.Printf("%v\t%T\tlength: %v\n", fifth, fifth, len(fifth))
	fmt.Printf("%v\t%T\tlength: %v\n", sixth, sixth, len(sixth))
	fmt.Printf("%v\t%T\tlength: %v\n", seventh, seventh, len(seventh))
}

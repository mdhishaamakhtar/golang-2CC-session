package main

import (
	"fmt"
)

func main() {
	first := make(map[string]int)
	first = map[string]int{
		"hishaam":  19,
		"hasan":    15,
		"sharanya": 18,
		"mayank":   19,
	}
	delete(first, "mayank")
	val1, ok1 := first["hishaam"]
	val2, ok2 := first["mayank"]
	fmt.Printf("%v\t%T\tlength: %v\n", first, first, len(first))
	fmt.Printf("%v\t%v\n", val1, ok1)
	fmt.Printf("%v\t%v\n", val2, ok2)
}

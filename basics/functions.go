package main

import (
	"fmt"
	"log"
)

func main() {
	var (
		str string
		i   int
		j   int
	)
	if _, err := fmt.Scanf("%v %v %v", &str, &i, &j); err != nil {
		log.Fatal(err)
	}
	SayMessage(str)
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(demonstratePointersPass(&i, &j))

	for i := 1; i < 10; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func SayMessage(msg string) {
	fmt.Printf("The message is %v\n", msg)
}

func sum(values ...int) int {
	fmt.Println(values)
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func demonstratePointersPass(i *int, j *int) int {
	*i = *i * 2
	return *i + *j
}

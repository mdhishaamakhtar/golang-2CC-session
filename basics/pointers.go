package main

import "fmt"

type Foo struct {
	foo int
	fii int
}

func main() {
	a := 42
	b := &a
	fmt.Printf("%v %p\n", a, b)

	c := &Foo{
		foo: 10,
		fii: 11,
	}
	fmt.Println(c)

	d := new(Foo)
	(*d).foo = 10
	d.fii = 11
	fmt.Println(d)
}

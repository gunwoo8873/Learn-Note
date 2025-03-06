package main

import (
	"fmt"
)

type Object struct {
	a, b, c string
}

func copy() Object {
	return Object{
		a: "a",
		b: "b",
		c: "c",
	}
}

func stack() {
	obj := copy()
	fmt.Println(obj)
}

func heap() Object {
	obj := copy() // Clone to stack memory in the value
	return obj
}

func main() {
	stack()
	fmt.Println(heap())
}

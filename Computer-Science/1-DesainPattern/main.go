package main

import (
	"fmt"
	"sync"
)

type Obj struct { // Create Obj datatype structure
	a int
}

var (
	instance *Obj
	once     sync.Once
)

func GetInstance() *Obj {
	once.Do(func() {
		instance = &Obj{a: 1}
	})

	return instance
}

func main() {
	objA := GetInstance()
	objB := GetInstance()

	fmt.Println(objA == objB)
}

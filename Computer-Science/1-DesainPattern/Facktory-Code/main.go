package main

import "fmt"

type Obj struct {
	name string
}

func (o *Obj) GetMethod() string {
	return o.name
}

func main() {
	obj := Obj{name: "User"}
	fmt.Println(obj.GetMethod())
}

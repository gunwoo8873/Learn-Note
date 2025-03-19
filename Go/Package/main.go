package main

import (
	"fmt"
	htemplate "html/template"
	"math/rand"

	"Go/Package/usepkg/custompkg" // Module in the package load
)

func main() {
	fmt.Println(rand.Int()) // Integer random value to -N ~ N [rand.Datatype<int, uint, float>]

	htemplate.New("test").Parse(`{{define "T"}}Hello`)

	custompkg.PrintCustom()
}

// Package Caluse : File in the Package name
package main

// Import to Package list and Module name is file name
import (
	"creat/bin/array"
	"creat/bin/control"
	"creat/bin/defined"
	"fmt"
)

// Module file in the func name is function call
func main() {
	array.Array()
	array.Slice()
	array.Map()

	defined.BindStructField()

	control.Container()
	control.Pointer()
	control.Switch_Case()

	fmt.Println("Hello World")
}

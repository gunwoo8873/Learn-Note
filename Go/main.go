// Package Caluse : File in the Package name
package main

// Import to Package list and Module name is file name
import (
	"creat/bin/array"
	"creat/bin/defined"
	"fmt"
)

// Module file in the func name is function call
func main() {
	array.Array()
	array.Slice()
	array.Map()

	defined.BindStructField()

	fmt.Println("Hello World")
}

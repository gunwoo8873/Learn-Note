// Package Caluse : File in the Package name
package main

// Import to Package list
import (
	"creat/helloworld"
	"creat/stream"
	"fmt"
)

func main() {
	helloworld.Hello()

	result := stream.Send()
	println(result)

	fmt.Println("Hello World")
}

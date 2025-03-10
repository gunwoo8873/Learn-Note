package main

import (
	"fmt"
)

func main() {
	result, err := myFunction()

	if err != nil {
		fmt.Println("Received an error: %s", err)
		return err
	}
}

package main

import (
	"fmt"
	"time"
)

var start = time.Now()

type Datatype struct {
	name        string
	email       string
	country     string
	phoneNumber string
}

func main() {
	p := Datatype{"John", "john@github.com", "USA", "1234567890"}
	fmt.Println(p)

	compile_time := time.Since(start)
	fmt.Printf("Compile time: %s \n", compile_time)
}

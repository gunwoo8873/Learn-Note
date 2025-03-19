package main

import (
	"fmt"
	"math/rand"
	"time"
)

/////////////////////////////////////////////
// Random int game
/////////////////////////////////////////////
// Object
// 1. Program compile start
// 2. Console in the respone number input
// 3. Seed number at input number equle

// "math/rand"
// "time"
// "os"
// "bufio"

func main() {
	setTime := time.Now().UnixNano()

	rand.Seed(setTime) // Set random time [UnixNano]

	n := rand.Intn(100)
	fmt.Println(n)
}

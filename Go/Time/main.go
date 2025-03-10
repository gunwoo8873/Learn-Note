package main

import (
	"fmt"
	"time"
)

func nowTime() {
	print := fmt.Println
	now := time.Now()
	print(now)
}

func fileterTime() {
	print := fmt.Println
	today := time.Now()
	print(today.Year())
	print(today.Month())
	print(today.Day())
	print(today.Hour())
	print(today.Minute())
	print(today.Second())
	print(today.Nanosecond())
	print(today.Location())
}

func main() {
	nowTime()
	fileterTime()
}

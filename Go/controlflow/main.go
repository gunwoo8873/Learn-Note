package main

import "fmt"

func ForLoop() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		}
		println(i)
	}
}

func ArrayLoop() {
	var arr = [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}

func MapLoop() {
	hashMap := map[int]string{1: "one", 2: "two", 3: "three"}
	for k, v := range hashMap {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
}

func main() {
	ForLoop()
	ArrayLoop()
	MapLoop()
}

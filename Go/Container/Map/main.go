package main

import (
	"fmt"
	"time"
)

var start = time.Now()

func createMap() {
	var Map = make(map[string]string) // make(map[key_type]value_type)
	Map["user1"] = "user1.github.com"
	Map["user2"] = "user2.github.com"

	for key, value := range Map {
		fmt.Printf("Static Map length: %d, Key: %v, Value: %v\n", len(Map), key, value)
	}
}

func deleteMap() {
	var Map = map[string]string{
		"user1.github.com": "192.168.33.15/3000",
		"user2.github.com": "192.168.33.16/3001",
	}
	for key, value := range Map {
		fmt.Printf("Static Map Key: %v, Value: %v\n", key, value)
	}

	var deleteUser = "user1.github.com"
	if _, exists := Map[deleteUser]; exists {
		delete(Map, deleteUser)
		fmt.Printf("Deleted success Map Key: %v \n", deleteUser)
	}

	for key, value := range Map {
		fmt.Printf("Has Map Key: %v, Value: %v\n", key, value)
	}
}

func main() {
	createMap()
	deleteMap()

	compile_time := time.Since(start)
	fmt.Printf("Compile time: %s \n", compile_time)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var start = time.Now()

func createMap() {
	var Map = make(map[string]string) // make(map[key_type]value_type)
	Map["user1"] = "user1.github.com"
	Map["user2"] = "user2.github.com"

	for key, value := range Map {
		fmt.Printf("Static Map Key: %v, Value: %v\n", key, value)
	}

	fmt.Printf("\n")
}

func updateMap() {
	var Map = map[string]string{
		"user1.linkedin.com": "127.0.0.1/3000",
	}

	newMap := Map
	newMap["user2.linkedin.com"] = "127.0.0.1/3001"

	for key, value := range Map {
		fmt.Printf("Update Map Key: %v, Value: %v\n", key, value)
	}

	fmt.Printf("\n")
}

func deleteMap() {
	var Map = map[string]string{
		"user1.github.com": "192.168.33.15/4000",
		"user2.github.com": "192.168.33.16/4001",
	}

	var deleteUser = "user1.github.com"
	if _, exists := Map[deleteUser]; exists {
		delete(Map, deleteUser)
		fmt.Printf("Deleted success Map Key: %v \n", deleteUser)
	}

	for key, value := range Map {
		fmt.Printf("Has Map Key: %v, Value: %v\n", key, value)
	}

	fmt.Printf("\n")
}

func syncMap() {
	var Map sync.Map

	// Store is map index save feature
	Map.Store("user1.outlook.kr", "132.115.44.11/5000")
	Map.Store("user2.outlook.kr", "132.115.44.12/5001")

	// sync.Map is repeat method use
	Map.Range(func(key, value any) bool {
		fmt.Printf("Sync Map Key: %v, Value: %v\n", key, value)
		return true
	})

	fmt.Printf("\n")
}

func main() {
	createMap()
	updateMap()
	deleteMap()
	syncMap()

	compile_time := time.Since(start)
	fmt.Printf("Compile time: %s \n", compile_time)
}

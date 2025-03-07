package main

import "fmt"

func Map() {
	dc := make(map[string]string)
	dc["user1"] = "user1.github.com"
	dc["user2"] = "user2.github.com"
	fmt.Printf("dc: length: %d, %v\n", len(dc), dc)

	inv := map[string]string{
		"user1.github.com": "192.168.33.15/3000",
		"user2.github.com": "192.168.33.16/3001",
	}
	fmt.Printf("inventory: length: %d, %v\n", len(inv), inv)

	delete(inv, "user1.github.com")
	fmt.Printf("inventory: length: %d, %v\n", len(inv), inv)
}

func main() {
	Map()
}

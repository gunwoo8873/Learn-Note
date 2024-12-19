package array

import "fmt"

func Map() {
	m := make(map[string]string)
	m["spine"] = "192.168.159.1" // Value bind to Spine label

	ip := m["spine"]
	ip, exists := m["spine"]

	if exists {
		fmt.Println(ip)
	}
}

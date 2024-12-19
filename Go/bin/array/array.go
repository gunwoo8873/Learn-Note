package array

import "fmt"

func Array() {
	hostname := [2]string{"localhost", "127.0.0.1"}

	fmt.Println(hostname[1])   // Array Index Element Select to Output
	fmt.Println(len(hostname)) // Array Index Element Count = 2

	ips := [3]string{
		"127.0.0.1/32",
		"192.168.0.1/32",
		"203.0.113.1/32",
	}

	fmt.Println(ips[2])
}

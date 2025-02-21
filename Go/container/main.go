package main

import "fmt"

func Array() {
	hostnames := []string{"user1.github.com", "user2.github.com"}
	for i := 0; i < len(hostnames); i++ {
		fmt.Println(hostnames[i])
	}
}

func Array_example() {
	var ipAddr [4]byte
	var localhost = [4]byte{127, 0, 0, 1}
	fmt.Println(len(localhost))   // 4
	fmt.Printf("%b\n", localhost) // {1111111}, {0}, {0}, {1}

	fmt.Println(ipAddr == localhost)
}

func Slice() {
	empty := []int{}
	fmt.Printf("empty: length: %d, capacity: %d, %v\n", len(empty), cap(empty), empty)

	words := []int{0, 1, 2, 3, -1, -2, -3}
	fmt.Printf("words: length: %d, capacity: %d, %v\n", len(words), cap(words), words)

	add := make([]int, 4)
	fmt.Printf("add: length: %d, capacity: %d, %v\n", len(add), cap(add), add)

	psuh := append(add, 5) // append = push or insert
	fmt.Printf("push: length: %d, capacity: %d, %v\n", len(psuh), cap(psuh), psuh)
}

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

func Collection() {
	type Router struct {
		Hostname  string
		IpAddr    string
		Platform  string
		Username  string
		Password  string
		Strictkey bool
	}

	type Inventory struct {
		Routers []Router
	}

	var user1 Router
	user1.Hostname = "http://promehteus.github.com"
	user1.IpAddr = "89.207.132.170"
	user1.Platform = "linux"
	user1.Username = "user1"
	user1.Password = "password1"
	user1.Strictkey = true

	user2 := Router{
		Hostname:  "http://promehteus.github.com",
		IpAddr:    "89.207.132.170",
		Platform:  "linux",
		Username:  "user2",
		Password:  "password2",
		Strictkey: true,
	}

	inv := Inventory{
		Routers: []Router{user1, user2},
	}
	fmt.Printf("inventory: %+v\n", inv)
}

func main() {
	Array()
	Array_example()
	Slice()
	Map()
	Collection()
}

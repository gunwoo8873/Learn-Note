package main

import "fmt"

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
	Collection()
}

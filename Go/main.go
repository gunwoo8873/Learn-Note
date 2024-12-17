// Package Caluse : File in the Package name
package main

// Import to Package list and Module name is file name
import (
	"lib/config"
	"lib/stream"
)

// Module file in the func name is function call
func main() {
	user := config.NewUser("Lee", "Warrior", "asd", "asd1234", "Lee@google.com")
	user.Userinfo()

	result := stream.Send()
	println(result)
}

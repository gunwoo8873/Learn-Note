package main

import (
	"fmt"
	"strings"
)

func genName(base string, suffix string) string {
	parts := []string{base, suffix}
	return strings.Join(parts, "-")
}

func processName(genName func(string, string) string, ip string) {
	base := "device"
	name := genName(base, ip)
	fmt.Println(name)
}

func main() {
	s := genName("device", "01")
	fmt.Println(s)

	processName(genName, "127.0.0.1")
}

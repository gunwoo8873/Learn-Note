package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Router struct {
	Hostname string `json:"hostname"`
	Ip       string `json:"ip"`
	ASN      uint16 `json:"asn"`
}

type Inventory struct {
	Router []Router `json:"router"`
}

func Encoding() {
	file, err := os.Open("./Processing/Json/sample.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()
	d := json.NewDecoder(file)

	var inv Inventory
	err = d.Decode(&inv)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("%+v\n", inv)
}

func main() {
	Encoding()
}

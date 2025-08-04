package main

import (
	"fmt"
	"math"
)

func integer() {
	int8 := math.MaxInt8
	fmt.Println("Int 8 bit :", int8)

	int16 := math.MaxInt16
	fmt.Println("Int 16 bit :", int16)

	int32 := math.MaxInt32 // Or uint8 name to 'rune' is UTF-8 Char data output
	fmt.Println("Int 32 bit :", int32)

	int64 := math.MaxInt64
	fmt.Println("Int 64 bit :", int64)
}

func uinteger() {
	uint8 := math.MaxUint8 // Or uint8 name to 'byte' is 1 byte data output
	fmt.Println("Uint 8 bit :", uint8)

	uint16 := math.MaxUint16
	fmt.Println("Uint 16 bit :", uint16)

	uint32 := math.MaxUint32
	fmt.Println("Uint 32 bit :", uint32)

	//// Short keyword using for variable to Int64 is max value limit error then add 'var' keyword use
	var uint64 uint64 = math.MaxUint64
	fmt.Println("Uint 64 bit :", uint64)
}

func float() {
	float32 := math.MaxFloat32
	fmt.Println("Float 32 bit :", float32)

	float64 := math.MaxFloat64
	fmt.Println("Float 64 bit :", float64)
}

func boolean() {
	bool := true
	fmt.Println("Boolean true :", bool)

	bool = false
	fmt.Println("Boolean false :", bool)
}

func main() {
	integer()
	uinteger()
	float()
	boolean()
}

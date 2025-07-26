package main

import (
	"fmt"
	"math"
)

func integer() {
	int8 := math.MaxInt8
	fmt.Println(int8)

	int16 := math.MaxInt16
	fmt.Println(int16)

	int32 := math.MaxInt32 // Or uint8 name to 'rune' is UTF-8 Char data output
	fmt.Println(int32)

	int64 := math.MaxInt64
	fmt.Println(int64)
}

func uinteger() {
	uint8 := math.MaxUint8 // Or uint8 name to 'byte' is 1 byte data output
	fmt.Println(uint8)

	uint16 := math.MaxUint16
	fmt.Println(uint16)

	uint32 := math.MaxUint32
	fmt.Println(uint32)

	//// Short keyword using for variable to Int64 is max value limit error then add 'var' keyword use
	var uint64 uint64 = math.MaxUint64
	fmt.Println(uint64)
}

func float() {
	float32 := math.MaxFloat32
	fmt.Println(float32)

	float64 := math.MaxFloat64
	fmt.Println(float64)
}

func boolean() {
	bool := true
	fmt.Println(bool)

	bool = false
	fmt.Println(bool)
}

func main() {
	integer()
	uinteger()
	float()
	boolean()
}

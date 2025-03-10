package main

import "fmt"

func tcpheader_example() {
	var headerWords uint8 = 5         // Header length is 5 words
	headerLen := headerWords * 32 / 8 // header len is 20bytes

	tcp_header := make([]byte, headerLen) // TCP header saved in 20bytes slice
	left_shift := headerLen << 4          // TCP header length equals is left shift 4 bits

	if len(tcp_header) > 13 {
		tcp_header[13] = tcp_header[13] | left_shift
		fmt.Printf("Updated TCP header: %08b\n", tcp_header[13])
	} else {
		fmt.Println("Error: tcp_header index out of bounds")
	}

	var tcpSyn uint8 = 1
	tcp_flags := tcpSyn << 1 // TCP SYN flas is left shift 1 bit

	if len(tcp_header) > 14 {
		tcp_header[14] = tcp_header[14] | tcp_flags
		fmt.Printf("Updated TCP header: %08b\n", tcp_header[14])
	} else {
		fmt.Println("Error: tcp_header index out of bounds")
	}

	tcpSynFlag := (tcp_header[14] & 0x02) != 0 // Check if TCP Syn flag is set
	parsedHeaderWords := tcp_header[13] >> 4
	fmt.Printf("TCP SYN flag: %t\nParsed header length: %d\n", tcpSynFlag, parsedHeaderWords)
}

func arithmetic() {
	var a, b int = 15, 40

	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	fmt.Printf("Sum: %d, Difference: %d, Product: %d, Quotient: %d, Remainder: %d\n", sum, diff, product, quotient, remainder)
}

func logical() {
	var a, b int = 4, 8

	and := a & b
	or := a | b
	xor := a ^ b
	not := ^a

	fmt.Printf("And: %d, Or: %d, Xor: %d, Not: %d\n", and, or, xor, not)

	equal := a == b
	notEqual := a != b
	lessThan := a < b
	lessThanEqual := a <= b
	greaterThan := a > b
	greaterThanEqual := a >= b

	fmt.Printf("Equal: %t, Not Equal: %t, Less Than: %t, Less Than Equal: %t, Greater Than: %t, Greater Than Equal: %t\n", equal, notEqual, lessThan, lessThanEqual, greaterThan, greaterThanEqual)
}

func comparison() {
	var a, b int = 4, 2

	left_shift := a << b
	right_shift := a >> b

	fmt.Printf("Left Shift: %d, Right Shift: %d\n", left_shift, right_shift)
}

func main() {
	tcpheader_example()
	arithmetic()
	logical()
	comparison()
}

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

	fmt.Printf(
		"Sum: %d\nDifference: %d\nProduct: %d\nQuotient: %d\nRemainder: %d\n",
		sum, diff, product, quotient, remainder,
	)
}

func logical() {
	var a, b int = 4, 8

	and := a & b // OR a && b 1 & 1 = 1, 0 & 1 = 0
	or := a | b  // OR a || b 1 | 1 = 1, 0 | 1 = 1, 0 | 0 = 0
	xor := a ^ b // 1 ^ 1 = 0, 0 ^ 1 = 1, 0 ^ 0 = 0
	not := ^a    // OR !a

	fmt.Printf(
		"And: %d, Or: %d, Xor: %d, Not: %d\n",
		and, or, xor, not,
	)

	equal := a == b
	notEqual := a != b
	lessThan := a < b
	lessThanEqual := a <= b
	greaterThan := a > b
	greaterThanEqual := a >= b

	fmt.Printf(
		"Equal: %t\nNot Equal: %t\nLess Than: %t\nLess Than Equal: %t\nGreater Than: %t\nGreater Than Equal: %t\n",
		equal, notEqual, lessThan, lessThanEqual, greaterThan, greaterThanEqual,
	)

	///////////////////////////////////////
	// || 										[1순위 연산자]
	// && 										[2순위 연산자]
	// ==, !=, <, <=, >, >= 	[3순위 연산자]
	// +, -, |, ^ 						[4순위 연산자]
	// *, /, %, <<, >>, &, &^ [5순위 연산자]
	///////////////////////////////////////
}

func stringLogical() {
	stra := "var"
	strb := "const"
	strc := "var"

	fmt.Printf("%s == %s : %v\n", stra, strb, stra == strb)
	fmt.Printf("%s != %s : %v\n", stra, strb, stra != strb)
	fmt.Printf("%s == %s : %v\n", stra, strc, stra == strc)
	fmt.Printf("%s != %s : %v\n", stra, strc, stra != strc)
}

func comparison() {
	var a int8 = 4

	a_left_shift := a << 2
	a_right_shift := a >> 2
	fmt.Printf("[8]Bit Left Shift: %08b\n[8]Right Shift: %08b\n", a_left_shift, a_right_shift)

	var b int8 = 64

	b_left_shift := b << 2
	b_right_shift := b >> 2
	fmt.Printf("[64]Bit Left Shift: %08b\n[64]Right Shift: %08b\n", b_left_shift, b_right_shift)
}

func main() {
	tcpheader_example()
	arithmetic()
	logical()
	stringLogical()
	comparison()
}

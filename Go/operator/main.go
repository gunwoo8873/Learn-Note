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

func main() {
	tcpheader_example()
}

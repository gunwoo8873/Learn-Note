package main

import "fmt"

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
	comparison()
}

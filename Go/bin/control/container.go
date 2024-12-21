package control

import (
	"fmt"
	"strings"
)

func Get_Resource(base string, suffix string) string {
	parts := []string{base, suffix}
	return strings.Join(parts, "-")
}

func Container() {
	s := Get_Resource("device", "01")
	fmt.Println(s)
}

type Device struct {
	name string
}

func mut(input *Device) {
	input.name += "-suffix"
}

func Pointer() {
	d := Device{name: "device"}
	mut(&d)
	fmt.Println(d.name)
}

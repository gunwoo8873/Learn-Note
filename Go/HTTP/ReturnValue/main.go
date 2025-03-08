package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

func getName(red io.Reader, wrt io.Writer) (string, error) {
	msg := "Test IO"
	fmt.Fprintf(wrt, msg)

	scanner := bufio.NewScanner(red)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("Empty name")
	}

	return name, nil
}

func main() {
	getName(strings.NewReader("Test"), nil)
}

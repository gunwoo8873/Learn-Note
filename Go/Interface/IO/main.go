package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	src := strings.NewReader("Copy text")
	dst, err := os.Create("./Interface/IO/create.txt") // './Interface/IO/create.txt'파일을 생성

	if err != nil {
		panic(err)
	}

	defer dst.Close()
	io.Copy(dst, src) // './Interface/IO/create.txt' 복사 후 덮어쓰기
}

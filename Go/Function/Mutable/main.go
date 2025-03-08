package main

import "fmt"

func addObject(object ...string) {
	fmt.Println(object.Join(object, "."))
}

func main() {

}

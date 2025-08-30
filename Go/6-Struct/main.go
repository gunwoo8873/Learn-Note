package main

/// Structure Field
type typeName struct {
	structA string
	structB int
}

func main() {
	var dataType typeName

	dataType.structA = "Hello World"
	dataType.structB = 10

	println(dataType.structA)
	println(dataType.structB)
}

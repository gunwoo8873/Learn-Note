package main

func main() {
	factory := Factory{}
	objA := factory.GetInstance()
	objB := factory.GetInstance()

	objA.Print()
	objB.Print()
}

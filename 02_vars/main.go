package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int 64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for unit8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name = "Carlos"
	var age = 30
	// const
	const isCool = true

	// shorthand
	lastName := "Alfaro"
	size := 1.3

	// More than one at the time
	fullName, email := "Carlos Alfaro", "carlos@cool.com"

	fmt.Println(name, age, isCool, lastName, size, fullName, email)
	fmt.Printf("%T\n", isCool)
}

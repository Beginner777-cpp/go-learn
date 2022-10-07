package main

import (
	"fmt"
)

func main() {
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// var name string = "test"
	const age int32 = 37

	//Shorthand
	// name := "Husan"
	name, email := "Husan", "abdigafurovhusan@gmail.com"
	size := 1.3

	fmt.Println(name, email, age)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T", size)
}

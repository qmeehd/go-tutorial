package main

import "fmt"

func main() {
	//bool

	//string

	//int  int8  int16  int32  int64
	//uint uint8 uint16 uint32 uint64 uintptr

	//byte // alias for uint8

	//rune // alias for int32
	// represents a Unicode code point

	//float32 float64

	//complex64 complex128
	var name string = "Jordan,"
	var age int = 25
	var isCool bool = true
	var email string = "jordan@gmail.com"

	//Shorthand
	//name := "Jordan"
	// ^ Detects that it's a string, works for var and int as well.

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", age)
}

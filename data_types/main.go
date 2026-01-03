package main

import (
	"fmt"
)

func main() {
	// bool
	var b bool = true // true or false
	fmt.Println(b)

	// integers (signed)
	var x int // based on your machine (32/64)
	var x int8
	var x int16
	var x int32 // "rune"
	var x int64

	// integers (unsigned)
	var x uint
	var x uint8 // byte
	var x uint16
	var x uint32
	var x uint64
	var x uintptr // for pointer arithmetic

	// float
	var x float32
	var x float64

	// complex
	var x complex64
	var x complex128

	// string
	var s string
	s = "hello" // immutable sequence of bytes
	// indexing a string s[i] will give you a byte (uint8)

	// array
	var a [3]int        // length is part of type
	a = [3]int{1, 2, 3} // array is fixed size, not dynamic

	// struct
	type Person struct {
		Name string
		Age  int
	}

	var p Person
	p.Name = "Duong"
	p.Age = 21

	// slice

	// map

	// pointer

	// function

	// channel

	// interface

	//
}

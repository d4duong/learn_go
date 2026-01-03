package main

import (
	"fmt"
)

/*
var — explicit declaration (works everywhere)
	- used at package level or inside functions
	- can specify type or let Go infer it
	- if don't assign a value, it gets assigned zero value (0, "", nil, etc)

:= — short declaration (function scope only)
	- used only inside functions (not at package level)
	- must have at least one new variable on left side
	- always declares and initializes in one step (no zero-value-nly form)

= — assignment to already declared variable

*/

// var
var x int       // x = 0
var s = "hello" // declaration of s with var, assignment with =, type inferred as string
var n int = 10  // declaration of n with var, assignment with =, type declared as int
var count int

const Pi = 3.14 // declaration with const, assignment with =

func main() {

	// var (explicit)
	var t bool
	fmt.Println(t)

	// :=  (short)
	t, y := true, 12
	fmt.Println(t, y)

	// =   (assignment)
	count = 12 // since count is already declared at package level
	fmt.Println(count, x, s, n)
	var k int // declaration
	k = 7     // assignment using =
	k = y + 1 // reassignment
	fmt.Println(k)
	v := []int{}
	var d []int
	fmt.Println(v == nil)
	fmt.Println(d == nil)

}

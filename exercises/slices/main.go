package main

import (
	"fmt"
)

// func Pic(dx, dy int) [][]uint8 {
// 	pic := make([]uint8, dx)
// 	printSlice("pic", pic)
// }

func main() {
	// pic.Show(Pic)
	dx := 10
	dy := 10
	var pic [][]uint8 = make([][]uint8, dy)
	for y := range dy {
		pic[y] = make([]uint8, dx)
	}
	printSlice("pic", pic)
	for x := range dx {
		for y := range dy {
			pic[x][y] = uint8(x * y)
		}
	}
	printSlice("pic", pic)
}

func printSlice(s string, x [][]uint8) {
	fmt.Printf("%s len=%d cap=%d\n",
		s, len(x), cap(x))
	for y := range len(x) {
		fmt.Println(x[y])
	}
}

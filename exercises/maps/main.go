package main

import (
	"fmt"
)

func WordCount(s string) map[string]int {
	return map[string]int{s: len(s)}
}

func main() {
	// m := make(map[string]int)
	// words := []string{
	// 	"hello",
	// 	"world",
	// 	"hi",
	// 	"reallylongword",
	// }
	m := WordCount("hello")
	for i, v := range m {
		fmt.Println(i, v)
	}

}

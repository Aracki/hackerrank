package main

import (
	"fmt"
)

func main() {
	s1 := "AABABAABBA"

	fmt.Println(alternatingCharacters(s1))
}

func alternatingCharacters(s string) int32 {

	counter := 0
	for i := 0; i < len(s)-1; i++ {
		s1 := string(s[i])
		sNext := string(s[i+1])
		if s1 == sNext {
			counter++
		}
	}
	return int32(counter)
}

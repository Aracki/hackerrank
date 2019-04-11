package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "hello"
	s2 := "world"
	fmt.Println(twoStrings(s1, s2))
}

func twoStrings(s1 string, s2 string) string {

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if string(s1[i]) == string(s2[j]) {
				return "YES"
			}
		}
	}

	return "NO"
}

func twoStringsFaster(s1 string, s2 string) string {

	if strings.ContainsAny(s1, s2) {
		return "YES"
	}
	return "NO"
}

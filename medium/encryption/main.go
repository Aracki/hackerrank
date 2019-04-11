package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	//s := encryption("haveaniceday")
	s := encryption("if man was meant to stay on the ground god would have given us roots")
	fmt.Println(s)
}

func removeSpace(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func encryption(s string) string {

	s = removeSpace(s)
	sqrt := math.Sqrt(float64(len(s)))

	rows := int(sqrt)
	columns := rows + 1

	fmt.Println("len: ", len(s))
	fmt.Println("sqrt: ", sqrt)
	fmt.Println("rows: ", rows)
	fmt.Println("columns: ", columns)

	var encoded strings.Builder

	for i := 0; i < columns; i++ {
		for j := i; j < len(s); j += columns {
			encoded.WriteString(string(s[j]))
		}
		encoded.WriteString(" ")
	}

	return encoded.String()
}

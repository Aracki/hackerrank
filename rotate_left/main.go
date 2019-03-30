package main

import (
	"fmt"
)

func rotLeft(a []int32, d int32) []int32 {

	ln := int32(len(a))
	var x = make([]int32, ln)
	var i int32
	for i = 0; i < ln; i++ {
		if i + d >= ln {
			x[i] = a[i+d-ln]
		} else {
			x[i] = a[i+d]
		}
	}

	return x
}

func main() {
	a := []int32{1, 2, 3, 4, 5}
    x := int32(3)

	l := rotLeft(a, x)
    fmt.Printf("rotate %v by %d \n", a, x)
	fmt.Println("result:", l)
}

package main

import "fmt"

func main() {

	arr := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	absDiff := diagonalDifference(arr)

	fmt.Println(absDiff)
}

func diagonalDifference(arr [][]int32) int32 {

	var firstD int32
	var secD int32

	for key, value := range arr {
		firstD += value[key]
		secD += value[len(arr)-1-key]
	}

	dif := firstD - secD

	if dif < 0 {
		return - dif
	} else {
		return dif
	}
}

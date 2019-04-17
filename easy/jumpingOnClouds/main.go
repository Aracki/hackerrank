package main

import "fmt"

func main() {

	//clouds := []int32{0, 1, 0, 0, 0, 1, 0}

	//c2 := []int32{0, 0, 1, 0, 0, 1, 0}
	c2 := []int32{0, 0, 0, 1, 0, 0}

	fmt.Println("res:", jumpingOnClouds(c2))
}

func jumpingOnClouds(clouds []int32) int32 {

	var jumps int32

	currentPosition := 0

	for currentPosition < len(clouds)-1 {

		if currentPosition == len(clouds)-2 {
			jumps++
			break
		}

		if clouds[currentPosition+2] == 0 {
			jumps++
			currentPosition += 2
		} else if clouds[currentPosition+2] == 1 {
			jumps++
			currentPosition++
		}
	}

	return jumps
}

package main

import "fmt"

func main() {

	dungeons := []int{-5, 1, 3, 10, -30, -5, 5}
	minH := minHealth(dungeons)
	fmt.Println(minH)
}

func minHealth(dungeons []int) int {

	var health, biggestDiff int

	for _, d := range dungeons {
		health += d
		if health < biggestDiff {
			biggestDiff = health
		}
	}
	return -1*biggestDiff + 1
}

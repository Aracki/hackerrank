package main

import (
	"fmt"
	"sort"
)

// TODO terminated due to timeout :(
func main() {

	n := int32(5)
	c := []int32{0, 4}

	fmt.Println(flatlandSpaceStations(n, c))
}

func flatlandSpaceStations(n int32, c []int32) int32 {

	var allDists []int32

	for i := int32(0); i < n; i++ {
		var dists []int32
		for _, station := range c {
			dist := station - i
			if dist < 0 {
				dist = -dist
			}
			dists = append(dists, dist)
			fmt.Println(dist)
		}
		sort.Slice(dists, func(i, j int) bool {
			return dists[i] < dists[j]
		})

		allDists = append(allDists, dists[0])
	}
	sort.Slice(allDists, func(i, j int) bool {
		return allDists[i] < allDists[j]
	})
	if len(allDists) < 1 {
		return 0
	}
	return allDists[len(allDists)-1]
}

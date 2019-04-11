package main

import (
	"fmt"
	"sort"
)

func main() {
	all := [][]int32{{5, 1}, {4, 0}, {6, 1}, {2, 1}, {8, 0}}
	n := luckBalance(2, all)
	fmt.Println(n)
}

func luckBalance(k int32, contests [][]int32) int32 {
	var luckAvl int32

	var imp []int
	for i := 0; i < len(contests); i++ {
		if contests[i][1] == 1 {
			imp = append(imp, int(contests[i][0]))
		} else {
			luckAvl = luckAvl + contests[i][0]
		}
	}

	// sort important ones
	sort.Ints(imp)

	for i := len(imp) - 1; i >= 0; i-- {
		if k != 0 {
			luckAvl = luckAvl + int32(imp[i])
			k --
		} else {
			luckAvl = luckAvl - int32(imp[i])
		}
	}

	return int32(luckAvl)
}

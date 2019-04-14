package main

import (
	"fmt"
	"sort"
)

// TODO not solved! [edge cases!]
func main() {

	b := biggerIsGreater2("fedcbabcd")
	fmt.Println(b)
}

func biggerIsGreater(w string) string {

	r := []rune(w)

	for i := len(w) - 1; i >= 0; i-- {
		if w[i] > w[i-1] {
			big := r[i]
			r[i] = r[i-1]
			r[i-1] = big
			return string(r)
		}
	}

	return "no answer"
}

func biggerIsGreater2(w string) string {

	r := []rune(w)
	printRunes(r)

	sort.Slice(r[:], func(i, j int) bool {
		return r[i] < r[j]
	})

	var head rune

	for x := 0; x < len(r)-1; x++ {

		for i := 0; i < len(r)-1; i++ {

			if string(r[i]) == string(w[x]) {
				if string(r[i]) != string(r[i+1]) {
					head = r[i+1]
					r = append(r[:i+1], r[i+2:]...)
					headArr := []rune{head}
					r = append(headArr, r...)

					printRunes(r)
					return string(r)
				}
			}
		}
	}

	return "no answer"
}

func printRunes(r []rune) {
	for _, v := range r {
		fmt.Println(v)
	}
	fmt.Println("***")
}

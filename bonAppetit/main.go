package main

import "fmt"

func main() {

	bill := []int32{3, 10, 2, 9}
	k := int32(1)
	b := int32(12)
	bonAppetit(bill, k, b)
}

func bonAppetit(bill []int32, k int32, b int32) {

	var owe int32

	billWithoutAnna := append(bill[:k], bill[k+1:]...)
	var sumWithoutAnna int32
	for _, v := range billWithoutAnna {
		sumWithoutAnna += v
	}

	if sumWithoutAnna/2 == b {
		fmt.Println("Bon Appetit")
		return
	}

	owe = sumWithoutAnna/2 - b

	fmt.Printf("%d", -owe)
}

package main

import (
	"fmt"
	"math/big"
)

func main() {
	extraLongFactorials(1000)
}

func extraLongFactorials(n int32) {

	fac := big.NewInt(int64(n))

	for n > 1 {
		//fac = fac * (n - 1)
		fac = fac.Mul(fac, big.NewInt(int64(n)-1))
		n--
	}

	fmt.Println(fac)
}

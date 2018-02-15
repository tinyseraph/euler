package main

import (
	"fmt"
	"math"
)

func main() {

	/*	2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
		What is the sum of the digits of the number 2^1000?
	*/

	for i := 1; i <= 15; i += 1 {
		trg := math.Pow(2, float64(i))
		fmt.Printf("number: %v => %v\n", trg, digitsum(trg))
	}

	trg := math.Pow(2, float64(100))
	fmt.Printf("pow: %v %v %v\n", trg, trg/10.0, math.Abs(trg))
	fmt.Printf("number: %v => %v\n", trg, digitsum(trg))
}

func digitsum(dgt float64) int {
	ret := 0
	clc := math.Abs(dgt)

	for {
		if clc < 10 {
			ret += int(clc)
			break
		} else {
			addnumber := int(clc) % 10
			ret += addnumber
			clc = clc / 10
		}
	}
	return ret
}

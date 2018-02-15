package main

import (
	"fmt"
	"math"
)

func main() {

	/*	2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
		What is the sum of the digits of the number 2^1000?
	*/

	/* badsample
	for i := 1; i <= 30; i += 1 {
		trg := math.Pow(2, float64(i))
		fmt.Printf("number: %v => %v\n", trg, digitsum(trg))
	}
	trg := math.Pow(2, float64(100))
	fmt.Printf("number: %v => %v\n", trg, digitsum(trg))
	*/

	/* try sample */
	testdata := [][]int{{2, 2}, {2, 8}, {2, 10}, {2, 15}, {2, 100}, {2, 1000}}
	sums := 0
	for _, x := range testdata {
		powed_num := largedigitPower(x[0], x[1])
		sums = array_sum(powed_num)
		fmt.Printf(">> arg: %v \n", x)
		fmt.Printf("   arry: %v \n   sum: %v\n", powed_num, sums)
		sums = 0
	}
}

func largedigitPower(base int, expon int) []int {
	digts := make([]int, 350)
	tmp := 0
	digts[0] = 1
	for i := 1; i <= expon; i += 1 {
		for idx, v := range digts {
			tmp1 := int((v * base) / 10)
			digts[idx] = (v*base)%10 + tmp
			tmp = tmp1
		}
	}
	return digts
}

func array_sum(target []int) int {
	ret := 0
	for _, x := range target {
		ret += x
	}
	return ret
}

/* bad sample */
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

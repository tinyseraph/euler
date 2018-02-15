package main

import (
	"fmt"
)

/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are
3 + 3 + 5 + 4 + 4 = 19 letters used in total.
If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters
and 115 (one hundred and fifteen) contains 20 letters.
The use of "and" when writing out numbers is in compliance with British usage.
*/

func main() {

	trg := []int{10, 105, 125, 200, 211, 222}
	for _, x := range trg {
		fmt.Printf("%v:\t%v \n", x, number_reader(x))
	}
	return
}

func number_reader(num int) string {

	words_map := map[int]string{1: "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "fourty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}

	tmp := ""
	res := ""
	if num/1000 == 1 {
		res = res + "one thousand"
	}
	if num%100 == 0 {
		tmp = fmt.Sprintf("%s handred", words_map[int(num/100)])
		res = res + tmp
	} else if num/100 > 0 && num%100 > 0 {
		_, chk := words_map[num%100]
		if chk == true {
			tmp = fmt.Sprintf("%s handred and %s", words_map[int(num/100)], words_map[num%100])
		} else {
			//check_two_digit := num % 100
			reduce_one_digit := num%100 - num%10
			one_digit := num % 10
			tmp = fmt.Sprintf("%s handred and %s %s",
				words_map[int(num/100)],
				words_map[reduce_one_digit],
				words_map[one_digit])
		}
		res = res + tmp
	}
	//	nty := num / 10
	_, chk := words_map[num]
	if chk == true {
		tmp := fmt.Sprintf("%s", words_map[num])
		res = res + tmp
	}

	return res
}

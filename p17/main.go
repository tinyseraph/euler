package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
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
	//check_letters(5)
	//check_letters(115)
	//check_letters(342)
	check_letters(1000)
}

func test_sample() {
	trg := []int{10, 115, 125, 200, 211, 222, 342}
	for _, x := range trg {
		numletter := number_reader(x)
		numletter_ns := strings.Replace(number_reader(x), " ", "", -1)
		fmt.Printf("%v\n\t%v\n\t%v\n\tsize:  %d\n", x, numletter, numletter_ns, utf8.RuneCountInString(numletter_ns))
	}
}

func test_sample5() {
	ret := ""
	for i := 1; i <= 5; i += 1 {
		ret = ret + number_reader(i)
	}
	fmt.Printf("%v\n%d", ret, len(ret))
}

func check_letters(num int) {
	ret := ""
	for i := 1; i <= num; i += 1 {
		ret = ret + number_reader(i)
	}
	ret_ns := strings.Replace(ret, " ", "", -1)
	fmt.Printf("letter's size: %d", utf8.RuneCountInString(ret_ns))
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
		40: "forty",
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
		return res
	}
	if num%100 == 0 {
		tmp = fmt.Sprintf("%s hundred", words_map[int(num/100)])
		res = res + tmp
	} else if num/100 > 0 && num%100 > 0 {
		_, chk := words_map[num%100]
		if chk == true {
			tmp = fmt.Sprintf("%s hundred and %s", words_map[int(num/100)], words_map[num%100])
		} else {
			//check_two_digit := num % 100
			reduce_one_digit := num%100 - num%10
			one_digit := num % 10
			tmp = fmt.Sprintf("%s hundred and %s %s",
				words_map[int(num/100)],
				words_map[reduce_one_digit],
				words_map[one_digit])
		}
		res = res + tmp
	}
	//	nty := num / 10
	if num < 100 {
		_, chk := words_map[num]
		if chk == true {
			tmp := fmt.Sprintf("%s", words_map[num])
			res = res + tmp
		} else {
			reduce_one_digit := num%100 - num%10
			one_digit := num % 10
			tmp = fmt.Sprintf("%s %s",
				words_map[reduce_one_digit],
				words_map[one_digit])
			res = res + tmp
		}
	}

	return res
}

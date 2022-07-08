package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func myAtoi(s string) int {

	sign := ' '
	num_start := false
	head_index := -1
	tail_index := -1
	whitespace := ' '

	for sindex, schar := range s {
		if schar == whitespace {
			continue
		} else if schar == '-' || schar == '+' {
			if sign == ' ' {
				sign = schar
			} else if sign != ' ' { //+-  or  -+ not allow
				return 0
			}

		} else if !unicode.IsNumber([]rune(s)[sindex]) && !num_start { // !num_start is the same as num_start == false
			return 0
		} else if unicode.IsNumber(schar) { //schar is number
			if schar == '0' && !num_start {
				continue
			}

			if schar != '0' && !num_start {
				num_start = true
				head_index = sindex
			}

			if sindex+1 < len(s) {
				if !unicode.IsNumber([]rune(s)[sindex+1]) { //check the rune of s[sindex+1]
					tail_index = sindex + 1
					break
				}
			}
		}
	}

	if head_index == -1 {
		return 0
	} else if head_index >= 0 {
		if tail_index == -1 {
			tail_index = len(s)
		}

		s = s[head_index:tail_index]
		ans_num, err := strconv.Atoi(s)

		if err != nil {
			return 0
		} else {
			if sign == '-' {
				ans_num = -1 * ans_num
			}

			if ans_num > 2147483647 {
				return 2147483647
			} else if ans_num < -2147483648 {
				return -2147483648
			} else {
				return ans_num
			}
		}
	}
	return 0
}

func main() {
	ss := "+-12"
	fmt.Print(myAtoi(ss))
}

package main

import "unicode"

func myAtoi(s string) int {
	positive := '+'
	pos_index := -1
	negative := '-'
	neg_index := -1
	num_start := false
	head_index := -1
	tail_index := -1
	whitespace := ' '

	for sindex, schar := range s {
		if schar == whitespace {
			continue
		} else if schar == negative {
			neg_index = sindex
		} else if schar == positive {
			pos_index = sindex
		} else if unicode.IsNumber(schar) && num_start == false { //schar is number
			if schar == '0' && num_start == false {
				continue
			}

			if schar != '0' && num_start == false {
				num_start = true
				head_index = sindex
			}
		} else if unicode.IsNumber(schar) == false && num_start == true { //number not continuous
			tail_index = sindex
			num_start = false
		}
	}

	if head_index >= 0 {

	}

	return 0
}

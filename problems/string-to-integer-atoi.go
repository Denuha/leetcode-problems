package problems

import (
	"math"
)

/*
	https://leetcode.com/problems/string-to-integer-atoi/
*/

func myAtoi(s string) int {
	i, result, sign, l := 0, 0, 1, len(s)
	MAX, MIN := (1<<31)-1, 1<<31
	// skip spaces
	for i < l && s[i] == ' ' {
		i++
	}
	// check if first non-space value is a sign
	if i < l && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	// iterate over digits
	for i < l && s[i] >= '0' && s[i] <= '9' {
		d := int(s[i] - '0')
		if sign > 0 && result > (MAX-d)/10 || sign < 0 && result > ((MIN-d)/10) {
			if sign == 1 {
				return MAX
			} else {
				return -MIN
			}
		}
		result = result*10 + d
		i++
	}
	return result * sign
}

func myAtoi2(s string) int {
	charMap := map[rune]int{
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'0': 0,
	}

	i := len(s)
	if i == 0 {
		return 0
	}
	rank := 1
	result := 0
	sign := 1
	firstSign := true
	firstLetter := false
	firstDigit := false
	firstDot := false

	for i > 0 {
		symb := rune(s[i-1])
		if symb == '+' && firstSign {
			sign = 1
			firstSign = false
			i -= 1
			continue
		}

		if symb == '-' && firstSign {
			sign = -1
			firstSign = false
			i -= 1
			continue
		}

		if symb == '.' {
			if firstDot {
				return 0
			}
			result = 0
			rank = 1
			firstDigit = true
			i -= 1
			continue
		}

		if val, ok := charMap[rune(s[i-1])]; ok {
			result = result + val*rank
			rank *= 10
			firstDigit = true
		} else {
			if symb != ' ' && !firstLetter && firstDigit {
				firstLetter = true
				return 0
			} else {
				return result
			}
		}

		i -= 1
	}

	result = result * sign
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}

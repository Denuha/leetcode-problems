package problems

import (
	"math"
)

/*
	https://leetcode.com/problems/reverse-integer/
*/
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
	}

	x *= sign

	var reverseDigit int

	for x > 0 {
		lastDigit := x % 10
		reverseDigit = reverseDigit*10 + lastDigit

		x /= 10
	}

	if reverseDigit > int(math.MaxInt32) {
		return 0
	}
	return reverseDigit * sign
}

package problems

import (
	"math"
)

/*
	https://leetcode.com/problems/divide-two-integers/
*/
func divide(dividend int, divisor int) int {
	sign := map[bool]int{true: -1, false: 1}[(dividend < 0) != (divisor < 0)]

	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))
	if dividend == divisor {
		return sign
	}

	k := 0

	for dividend >= divisor {
		dividend -= divisor
		k += 1
	}

	if k*sign > int(math.MaxInt32) {
		k = math.MaxInt32
	}

	if k*sign < -int(math.MaxInt32) {
		k = 2147483648
	}
	return k * sign
}

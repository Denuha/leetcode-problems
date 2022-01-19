package problems

func DivideTest() (bool, string) {
	var (
		dividend int
		divisor  int
	)

	dividend = 10
	divisor = 3
	if divide(dividend, divisor) != 3 {
		return false, "case 1"
	}

	dividend = 7
	divisor = -3
	if divide(dividend, divisor) != -2 {
		return false, "case 2"
	}

	dividend = 2147483647
	divisor = 1
	if divide(dividend, divisor) != 2147483647 {
		return false, "case 3"
	}

	dividend = -2147483648
	divisor = -1
	if divide(dividend, divisor) != 2147483647 {
		return false, "case 4"
	}

	dividend = -2147483648
	divisor = 1
	if divide(dividend, divisor) != -2147483648 {
		return false, "case 5"
	}

	return true, "good"
}

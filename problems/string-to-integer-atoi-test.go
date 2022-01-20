package problems

func MyAtoiTest() (bool, string) {
	var s string

	s = "42"
	if myAtoi(s) != 42 {
		return false, "case 1"
	}

	s = "   -42"
	if myAtoi(s) != -42 {
		return false, "case 2"
	}

	s = "4193 with words"
	if myAtoi(s) != 4193 {
		return false, "case 3"
	}

	s = "words and 987"
	if myAtoi(s) != 0 {
		return false, "case 4"
	}

	s = "3.141592"
	if myAtoi(s) != 3 {
		return false, "case 5"
	}

	s = "  -0012a42"
	if myAtoi(s) != -12 {
		return false, "case 6"
	}

	s = "00000-42a1234"
	if myAtoi(s) != 0 {
		return false, "case 7"
	}

	s = "9223372036854775808"
	if myAtoi(s) != 2147483647 {
		return false, "case 8"
	}

	s = "-91283472332"
	if myAtoi(s) != -2147483648 {
		return false, "case 9"
	}

	s = "  0000000000012345678"
	if myAtoi(s) != 12345678 {
		return false, "case 10"
	}

	s = "123-"
	if myAtoi(s) != 123 {
		return false, "case 11"
	}

	s = "  +  413"
	if myAtoi(s) != 0 {
		return false, "case 12"
	}

	return true, "good"
}

package problems

func ReverseTest() (bool, string) {
	var x int

	x = 123
	if reverse(x) != 321 {
		return false, "case 1"
	}

	x = -123
	if reverse(x) != -321 {
		return false, "case 2"
	}

	x = 120
	if reverse(x) != 21 {
		return false, "case 3"
	}

	return true, "good"
}

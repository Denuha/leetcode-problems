package problems

func intToRoman(num int) string {
	digits := []int{1000, 900, 500, 400, 100, 90,
		50, 40, 10, 9, 5, 4, 1}

	romans := []string{"M", "CM", "D", "CD", "C", "XC",
		"L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for idx, dig := range digits {
		times := num / dig
		num %= dig
		for times > 0 {
			result += romans[idx]
			times--
		}
	}

	return result
}

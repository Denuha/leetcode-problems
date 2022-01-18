package problems

import "reflect"

func IntToRomanTest() (bool, string) {
	var num int

	num = 3
	if !reflect.DeepEqual(intToRoman(num), "III") {
		return false, "case 1"
	}

	num = 58
	if !reflect.DeepEqual(intToRoman(num), "LVIII") {
		return false, "case 2"
	}

	num = 1994
	if !reflect.DeepEqual(intToRoman(num), "MCMXCIV") {
		return false, "case 3"
	}

	return true, "good"
}

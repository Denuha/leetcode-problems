package problems

import "reflect"

func WordPatternTest() (bool, string) {
	var (
		pattern string
		s       string
	)

	pattern = "abba"
	s = "dog cat cat dog"
	if !reflect.DeepEqual(wordPattern(pattern, s), true) {
		return false, "case 1"
	}

	pattern = "abba"
	s = "dog cat cat fish"
	if !reflect.DeepEqual(wordPattern(pattern, s), false) {
		return false, "case 2"
	}

	pattern = "aaaa"
	s = "dog cat cat dog"
	if !reflect.DeepEqual(wordPattern(pattern, s), false) {
		return false, "case 3"
	}

	return true, "good"
}

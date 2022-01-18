package problems

import "strings"

/*
	https://leetcode.com/problems/word-pattern/
*/
func wordPattern(pattern string, s string) bool {
	mapWord := make(map[string]rune)
	usedLetters := make([]bool, 26)
	var words = strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	for i, word := range words {
		if _, ok := mapWord[word]; !ok && !usedLetters[pattern[i]-'a'] {
			mapWord[word] = rune(pattern[i])
			usedLetters[pattern[i]-'a'] = true
		} else {
			currRune := mapWord[word]
			if currRune != rune(pattern[i]) {
				return false
			}
		}
	}

	return true
}

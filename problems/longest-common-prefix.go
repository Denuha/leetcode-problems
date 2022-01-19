package problems

import (
	"sort"
)

/*
	https://leetcode.com/problems/longest-common-prefix/
*/
func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]

	c := 0
	for c < len(first) {
		if first[c] == last[c] {
			c += 1
		} else {
			break
		}
	}

	if c == 0 {
		return ""
	} else {
		return first[:c]
	}
}

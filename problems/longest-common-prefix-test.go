package problems

func LongestCommonPrefixTest() (bool, string) {
	var input []string

	input = []string{"flower", "flow", "flight"}
	if longestCommonPrefix(input) != "fl" {
		return false, "case 1"
	}

	input = []string{"dog", "racecar", "car"}
	if longestCommonPrefix(input) != "" {
		return false, "case 2"
	}

	return true, "good"
}

package problems

/*
	https://leetcode.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	sumsMap := make(map[int]int)

	for idx, num := range nums {
		if i, ok := sumsMap[target-num]; ok {
			return []int{i, idx}
		}
		sumsMap[num] = idx
	}

	return nil
}

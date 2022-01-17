package problems

/*
	https://leetcode.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	var result []int

	contains := func(num int, nums map[int]int) bool {
		if _, ok := nums[num]; ok {
			return true
		}
		return false
	}

	indexOfArr := func(num int, arr []int, flagRepeat bool) int {
		for i, v := range arr {
			if num == v {
				if flagRepeat {
					flagRepeat = false
					continue
				}
				return i
			}
		}
		return -1
	}

	sumsMap := make(map[int]int)

	for _, num := range nums {
		sumsMap[num] = target - num
	}

	for key, value := range sumsMap {
		flagRepeat := false
		if key == value {
			flagRepeat = true
		}
		if contains(value, sumsMap) {
			a := indexOfArr(key, nums, false)
			b := indexOfArr(value, nums, flagRepeat)
			if a < 0 || b < 0 {
				continue
			}
			result = append(result, a, b)
			break
		}
	}

	return result
}

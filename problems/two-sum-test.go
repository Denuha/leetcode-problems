package problems

import "reflect"

func TwoSumTest() bool {
	var (
		nums   []int
		target int
	)

	nums = []int{2, 7, 11, 15}
	target = 9
	if !reflect.DeepEqual(twoSum(nums, target), []int{0, 1}) {
		return false
	}

	nums = []int{3, 2, 4}
	target = 6
	if !reflect.DeepEqual(twoSum(nums, target), []int{1, 2}) {
		return false
	}

	nums = []int{3, 3}
	target = 6
	if !reflect.DeepEqual(twoSum(nums, target), []int{0, 1}) {
		return false
	}

	return true
}

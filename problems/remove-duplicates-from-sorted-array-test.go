package problems

import (
	"reflect"
)

func RemoveDuplicatesTest() (bool, string) {
	var (
		nums         []int
		expectedNums []int
		k            int
	)

	nums = []int{1, 1, 2}
	expectedNums = []int{1, 2}
	k = removeDuplicates(nums)
	if k != len(expectedNums) {
		return false, "case 1"
	}

	if !reflect.DeepEqual(nums, expectedNums) {
		return false, "case 1"
	}

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedNums = []int{0, 1, 2, 3, 4}

	k = removeDuplicates(nums)
	if k != len(expectedNums) {
		return false, "case 2"
	}

	if !reflect.DeepEqual(nums, expectedNums) {
		return false, "case 2"
	}

	return true, "good"
}

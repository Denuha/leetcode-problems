package problems

import (
	"reflect"
)

func RemoveDuplicatesTest() bool {
	var (
		nums         []int
		expectedNums []int
		k            int
	)

	nums = []int{1, 1, 2}
	expectedNums = []int{1, 2}
	k = removeDuplicates(nums)
	if k != len(expectedNums) {
		return false
	}

	if !reflect.DeepEqual(nums, expectedNums) {
		return false
	}

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedNums = []int{0, 1, 2, 3, 4}

	k = removeDuplicates(nums)
	if k != len(expectedNums) {
		return false
	}

	if !reflect.DeepEqual(nums, expectedNums) {
		return false
	}

	return true
}

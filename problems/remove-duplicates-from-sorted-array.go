package problems

/*
	https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/

func removeDuplicates(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}

	j := 0 // points to  the index of last filled position
	for i := 1; i < ln; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1

	/*
		var numsMap = make(map[int]int)
		for i := 0; i < len(nums); i++ {
			num := nums[i]
			if _, ok := numsMap[num]; !ok {
				numsMap[num] = num
			} else {
				nums = append(nums[:i], nums[i+1:]...)
				i -= 1
			}
		}

		fmt.Println(nums)
		return len(nums)
	*/
}

package problems

import (
	"sort"
)

/*
	https://leetcode.com/problems/median-of-two-sorted-arrays
	O(m+n) - 1
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	length := len(nums1)

	if length%2 != 0 {
		return float64(nums1[length/2])
	} else {
		var a float64 = float64(nums1[length/2-1])
		var b float64 = float64(nums1[length/2])
		var c float64 = (a + b) / 2
		return c
	}
}

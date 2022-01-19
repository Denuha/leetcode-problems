package problems

func FindMedianSortedArraysTest() (bool, string) {
	var (
		nums1 []int
		nums2 []int
	)

	nums1 = []int{1, 3}
	nums2 = []int{2}
	if findMedianSortedArrays(nums1, nums2) != 2.0 {
		return false, "case 1"
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	if findMedianSortedArrays(nums1, nums2) != 2.5 {
		return false, "case 2"
	}

	return true, "good"
}

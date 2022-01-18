package problems

import "reflect"

func CanPlaceFlowersTest() (bool, string) {
	var (
		flowerbed []int
		n         int
	)

	flowerbed = []int{1, 0, 0, 0, 1}
	n = 1
	if !reflect.DeepEqual(canPlaceFlowers(flowerbed, n), true) {
		return false, "case 1"
	}

	flowerbed = []int{1, 0, 0, 0, 1}
	n = 2
	if !reflect.DeepEqual(canPlaceFlowers(flowerbed, n), false) {
		return false, "case 2"
	}

	return true, "good"
}

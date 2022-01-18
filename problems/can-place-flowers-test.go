package problems

import "reflect"

func CanPlaceFlowersTest() bool {
	var (
		flowerbed []int
		n         int
	)

	flowerbed = []int{1, 0, 0, 0, 1}
	n = 1
	if !reflect.DeepEqual(canPlaceFlowers(flowerbed, n), true) {
		return false
	}

	flowerbed = []int{1, 0, 0, 0, 1}
	n = 2
	if !reflect.DeepEqual(canPlaceFlowers(flowerbed, n), false) {
		return false
	}

	return true
}

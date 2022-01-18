package problems

import (
	"strconv"
	"strings"
)

/*
	https://leetcode.com/problems/can-place-flowers/
*/

func canPlaceFlowers(flowerbed []int, n int) bool {
	var flowerbedStr string

	for _, f := range flowerbed {
		flowerbedStr += strconv.Itoa(f)
	}

	flowerbedStr = "0" + flowerbedStr + "0"

	for n > 0 {
		if !strings.Contains(flowerbedStr, "000") {
			return false
		}
		flowerbedStr = strings.Replace(flowerbedStr, "000", "010", 1)
		n -= 1
	}
	return true
}

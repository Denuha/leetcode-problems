package main

import (
	"fmt"

	"github.com/Denuha/leetcode-problems/problems"
)

func main() {
	fmt.Println("twoSum:", problems.TwoSumTest())

	fmt.Println("Add Two Numbers:", problems.AddTwoNumbersTest())

	fmt.Println("Remove Duplicates from Sorted Array:", problems.RemoveDuplicatesTest())

	fmt.Println("Pow(x,n):", problems.MyPowTest())

	fmt.Println("Can Place Flowers:", problems.CanPlaceFlowersTest())

	ok, str := problems.WordPatternTest()
	fmt.Println("Word Pattern:", ok, str)
}

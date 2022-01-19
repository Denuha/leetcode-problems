package main

import (
	"fmt"

	"github.com/Denuha/leetcode-problems/problems"
)

func main() {
	var (
		ok         bool
		resultCase string
	)

	// ok, resultCase = problems.TwoSumTest()
	// fmt.Println("twoSum:", ok, resultCase)

	// ok, resultCase = problems.AddTwoNumbersTest()
	// fmt.Println("Add Two Numbers:", ok, resultCase)

	// ok, resultCase = problems.RemoveDuplicatesTest()
	// fmt.Println("Remove Duplicates from Sorted Array:", ok, resultCase)

	// ok, resultCase = problems.MyPowTest()
	// fmt.Println("Pow(x,n):", ok, resultCase)

	// ok, resultCase = problems.CanPlaceFlowersTest()
	// fmt.Println("Can Place Flowers:", ok, resultCase)

	// ok, resultCase = problems.WordPatternTest()
	// fmt.Println("Word Pattern:", ok, resultCase)

	// ok, resultCase = problems.IntToRomanTest()
	// fmt.Println("Integer to Roman:", ok, resultCase)

	// ok, resultCase = problems.ReverseTest()
	// fmt.Println("Reverse Integer:", ok, resultCase)

	// ok, resultCase = problems.DivideTest()
	// fmt.Println("Divide:", ok, resultCase)

	// ok, resultCase = problems.LongestCommonPrefixTest()
	// fmt.Println("Longest Common Prefix:", ok, resultCase)

	ok, resultCase = problems.FindMedianSortedArraysTest()
	fmt.Println("Median of Two Sorted Arrays:", ok, resultCase)
}

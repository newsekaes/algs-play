package main

import (
	"fmt"
)

func maxProfit(nums []int) int {
	len := len(nums)
	dpMin, dpMax := nums[0], 0
	for i := 1; i < len; i++ {
		dpMin = min(dpMin, nums[i])
		dpMax = max(nums[i]-dpMin, dpMax)
	}
	return dpMax
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("%d\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%d\n", maxProfit([]int{7, 6, 4, 3, 1}))
}

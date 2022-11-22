package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	len := len(nums)
	maxNum, dpMin, dpMax := nums[0], nums[0], nums[0]
	for i := 1; i < len; i++ {
		if nums[i] > 0 {
			dpMin, dpMax = min(nums[i], nums[i]*dpMin), max(nums[i], nums[i]*dpMax)
		} else {
			dpMin, dpMax = min(nums[i], nums[i]*dpMax), max(nums[i], nums[i]*dpMin)
		}
		maxNum = max(maxNum, dpMax)
	}
	return maxNum
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
	fmt.Printf("%d\n", maxProduct([]int{2, 3, -2, 4}))
	fmt.Printf("%d\n", maxProduct([]int{2, 3, 2, 4}))
	fmt.Printf("%d\n", maxProduct([]int{2}))
	fmt.Printf("%d\n", maxProduct([]int{-2, 0, -1}))
	fmt.Printf("%d\n", maxProduct([]int{-1, 1, 0, -1, 1, -1}))
	fmt.Printf("%d\n", maxProduct([]int{-1, 1, -1, 1, -1}))
}

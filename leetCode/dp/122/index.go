package main

import (
	"fmt"
)

func maxProfit(nums []int) int {
	len := len(nums)
	dpMin := nums[0]
	sale := 0
	for i := 1; i < len; i++ {
		if nums[i] > dpMin {
			sale += nums[i] - dpMin
			dpMin = nums[i]
		} else {
			dpMin = nums[i]
		}
	}
	return sale
}

func main() {
	fmt.Printf("%d\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%d\n", maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Printf("%d\n", maxProfit([]int{7, 6, 4, 3, 1}))
}

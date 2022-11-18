package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	level := len(triangle)
	dp := triangle[level-1]
	for i := level - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	in := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Printf("%d\n", minimumTotal(in))
}

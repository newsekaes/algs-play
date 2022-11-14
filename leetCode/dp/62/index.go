package main

import "fmt"

func uniquePaths(m int, n int) int {
	// 保证 n 是最小的
	if n > m {
		m, n = n, m
	}
	dp := make([]int, n)

	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				continue
			} else {
				dp[j] = dp[j-1] + dp[j]
			}
		}
	}

	return dp[len(dp)-1]
}

func main() {
	fmt.Print(uniquePaths(3, 3), "\n")
}

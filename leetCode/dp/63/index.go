package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 保证 n 是最小的
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	dp := make([]int, m)

	for i := range dp {
		if obstacleGrid[0][i] == 0 {
			dp[i] = 1
		} else {
			dp[i] = 0
			break
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 0 {
				if j == 0 {
					continue
				} else {
					dp[j] = dp[j-1] + dp[j]
				}
			} else {
				dp[j] = 0
			}

		}
	}

	return dp[len(dp)-1]
}

func main() {
	fmt.Print(uniquePathsWithObstacles([][]int{{1, 0}}), "\n")
}

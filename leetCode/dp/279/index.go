package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	f := make([]int, n+1)

	for i := 1; i <= n; i++ {
		mmin := math.MaxInt32 - 1
		for j := 1; j*j <= i; j++ {
			mmin = min(mmin, f[i-j*j])
			if mmin == 0 {
				break
			}
		}
		f[i] = mmin + 1
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Printf("%d\n", numSquares(i))
	}
	fmt.Printf("%d\n", numSquares(52))
}

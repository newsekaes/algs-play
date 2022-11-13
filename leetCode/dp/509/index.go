package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		var res int
		dp := [2]int{0, 1}
		for i := 2; i <= n; i++ {
			res = fib_help(&dp)
		}
		return res
	}
}

func fib_help(dp *[2]int) int {
	res := dp[0] + dp[1]
	dp[0] = dp[1]
	dp[1] = res
	return res
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d\n", fib(i))
	}
}

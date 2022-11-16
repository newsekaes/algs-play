package main

import "fmt"

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main() {
	for i := 1; i < 15; i++ {
		fmt.Printf("%d\n", climbStairs(i))
	}
}

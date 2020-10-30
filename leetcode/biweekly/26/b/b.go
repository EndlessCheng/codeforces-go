package main

import "fmt"

// github.com/EndlessCheng/codeforces-go
func simplifiedFractions(n int) (ans []string) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				ans = append(ans, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}
	return
}

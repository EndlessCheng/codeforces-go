package main

import "math/big"

// github.com/EndlessCheng/codeforces-go
func numPrimeArrangements(n int) int {
	c := 0
o:
	for v := 2; v <= n; v++ {
		for d := 2; d*d <= v; d++ {
			if v%d == 0 {
				continue o
			}
		}
		c++
	}
	ans := new(big.Int).MulRange(1, int64(c))
	return int(ans.Mul(ans, new(big.Int).MulRange(1, int64(n-c))).Rem(ans, big.NewInt(1e9+7)).Int64())
}

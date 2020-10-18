package main

import "math/big"

// github.com/EndlessCheng/codeforces-go
func numberOfSets(n int, k int) int {
	return int(new(big.Int).Mod(new(big.Int).Binomial(int64(n+k-1), int64(2*k)), big.NewInt(1e9+7)).Int64())
}

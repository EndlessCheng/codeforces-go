package main

import "math/big"

// github.com/EndlessCheng/codeforces-go
func numberOfWays(n int) int {
	return int(new(big.Int).Mod(new(big.Int).Div(new(big.Int).Binomial(int64(n), int64(n/2)), big.NewInt(int64(n/2+1))), big.NewInt(1e9+7)).Int64())
}

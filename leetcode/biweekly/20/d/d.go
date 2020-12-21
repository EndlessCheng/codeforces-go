package main

import "math/big"

// github.com/EndlessCheng/codeforces-go
func countOrders(n int) int {
	return int(new(big.Int).Mod(new(big.Int).Rsh(new(big.Int).MulRange(1, int64(2*n)), uint(n)), big.NewInt(1e9+7)).Int64())
}

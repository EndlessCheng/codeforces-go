package main

import "math/big"

// github.com/EndlessCheng/codeforces-go
func countVowelStrings(n int) int {
	return int(new(big.Int).Binomial(int64(n+4), 4).Int64())
}

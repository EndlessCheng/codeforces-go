package main
import ."math/big"

// github.com/EndlessCheng/codeforces-go
func solve_bangbang(n, m, k int) int64 {
	if m == 0 {
		return 1
	}
	n -= k * (m - 1)
	if n < m {
		return 0
	}
	return new(Int).Mod(new(Int).Binomial(int64(n), int64(m)), NewInt(1e9+7)).Int64()
}

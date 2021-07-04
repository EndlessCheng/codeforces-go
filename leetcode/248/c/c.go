package main

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7

func countGoodNumbers(n int64) int {
	return pow(5, (int(n)+1)/2) * pow(4, int(n)/2) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

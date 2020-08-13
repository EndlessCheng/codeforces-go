package main

// github.com/EndlessCheng/codeforces-go
func GetNumberOfSchemes(n int) int {
	n--
	const mod int = 1e9 + 7
	pow := func(x, n int) (res int) {
		x %= mod
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}
	return (pow(2, n) + pow(4, n)) % mod
}

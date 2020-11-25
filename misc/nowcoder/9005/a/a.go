package main

// github.com/EndlessCheng/codeforces-go
func tree4(N int64) int64 {
	const mod = 998244353
	n := int(N)
	ans, i := 0, 0
	for ; n >= 1<<i; i++ {
		st := 1 << i
		end := st<<1 - 1
		s := (st + end) * st / 2 % mod
		ans += s * (i + 1) % mod
		n -= 1 << i
	}
	if n > 0 {
		st := 1 << i
		end := st + n - 1
		s := (st + end) * n / 2 % mod
		ans += s * (i + 1) % mod
	}
	return int64(ans % mod)
}

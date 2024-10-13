package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 1001

var s [mx][mx]int

func init() {
	s[0][0] = 1
	for i := 1; i < mx; i++ {
		for j := 1; j <= i; j++ {
			s[i][j] = (s[i-1][j-1] + j*s[i-1][j]) % mod
		}
	}
}

func numberOfWays(n, x, y int) (ans int) {
	perm, powY := 1, 1
	for i := 1; i <= min(n, x); i++ {
		perm = perm * (x + 1 - i) % mod
		powY = powY * y % mod
		ans = (ans + perm*s[n][i]%mod*powY) % mod
	}
	return
}

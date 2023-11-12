package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007

func stringCount(n int) (ans int) {
	return ((pow(26, n)-
		     pow(25, n-1)*(75+n)+
		     pow(24, n-1)*(72+n*2)-
		     pow(23, n-1)*(23+n))%mod + mod) % mod // 保证结果非负
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

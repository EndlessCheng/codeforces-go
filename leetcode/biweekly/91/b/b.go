package main

// https://space.bilibili.com/206214
func countGoodStrings(low, high, zero, one int) (ans int) {
	g := gcd(zero, one)
	low = (low-1)/g + 1
	high /= g
	zero /= g
	one /= g

	const mod = 1_000_000_007
	f := make([]int, high+1) // f[i] 表示构造长为 i 的字符串的方案数
	f[0] = 1                 // 构造空串的方案数为 1
	for i := 1; i <= high; i++ {
		if i >= one {
			f[i] = f[i-one]
		}
		if i >= zero {
			f[i] = (f[i] + f[i-zero]) % mod
		}
		if i >= low {
			ans = (ans + f[i]) % mod
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

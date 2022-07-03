package main

// https://space.bilibili.com/206214/dynamic
func peopleAwareOfSecret(n, delay, forget int) int {
	const mod int = 1e9 + 7
	sum := make([]int, n+1)
	sum[1] = 1
	for i := 2; i <= n; i++ {
		f := (sum[max(0, i-delay)] - sum[max(0, i-forget)]) % mod
		sum[i] = (sum[i-1] + f) % mod
	}
	return ((sum[n]-sum[max(0, n-forget)])%mod + mod) % mod // 防止结果为负数
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

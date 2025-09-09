package main

// https://space.bilibili.com/206214/dynamic
func peopleAwareOfSecret(n, delay, forget int) int {
	const mod = 1_000_000_007
	sum := make([]int, n+1) // known 数组的前缀和
	sum[1] = 1

	for j := 2; j <= n; j++ {
		known := sum[max(j-delay, 0)] - sum[max(j-forget, 0)]
		sum[j] = (sum[j-1] + known) % mod
	}

	ans := sum[n] - sum[max(n-forget, 0)]
	return (ans%mod + mod) % mod // 保证答案非负
}

func peopleAwareOfSecret1(n, delay, forget int) (ans int) {
	const mod = 1_000_000_007
	diff := make([]int, n+2)
	diff[1] = 1
	diff[2] = -1
	known := 0

	for i := 1; i <= n; i++ {
		// 加上 diff[i] 后，known 表示恰好在第 i 天得知秘密的人数
		known = (known + diff[i]) % mod
		// 统计在第 n 天没有忘记秘密的人数
		if i >= n-forget+1 {
			ans += known
		}
		// 恰好在第 i 天得知秘密的人，会在第 [i+delay, i+forget-1] 天分享秘密
		diff[min(i+delay, n+1)] += known
		diff[min(i+forget, n+1)] -= known // 注意这里有减法，这会导致上面累加 diff[i] 时，known 可能是负数
	}

	return (ans%mod + mod) % mod // 保证答案非负
}

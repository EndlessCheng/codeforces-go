package main

// https://space.bilibili.com/206214
func findCoins(numWays []int) (ans []int) {
	n := len(numWays)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		ways := numWays[i-1]
		if ways == f[i] {
			continue
		}
		if ways-1 != f[i] {
			return nil
		}
		ans = append(ans, i)
		// 现在得到了一个大小为 i 的物品，用 i 计算完全背包
		for j := i; j <= n; j++ {
			f[j] += f[j-i]
		}
	}
	return
}

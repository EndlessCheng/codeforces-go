package main

// https://space.bilibili.com/206214
func waysToReachTarget(target int, types [][]int) int {
	const mod = 1_000_000_007
	f := make([]int, target+1)
	f[0] = 1
	for _, p := range types {
		count, marks := p[0], p[1]
		for j := target; j > 0; j-- {
			for k := 1; k <= min(count, j/marks); k++ {
				f[j] += f[j-k*marks]
			}
			f[j] %= mod
		}
	}
	return f[target]
}

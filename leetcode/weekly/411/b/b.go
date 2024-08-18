package main

// https://space.bilibili.com/206214
func maxEnergyBoost(a, b []int) int64 {
	n := len(a)
	f := make([][2]int64, n+2)
	for i, x := range a {
		f[i+2][0] = max(f[i+1][0], f[i][1]) + int64(x)
		f[i+2][1] = max(f[i+1][1], f[i][0]) + int64(b[i])
	}
	return max(f[n+1][0], f[n+1][1])
}

func maxEnergyBoost2(a, b []int) int64 {
	n := len(a)
	c := [2][]int{a, b}
	memo := make([][2]int64, n)
	var dfs func(int, int) int64
	dfs = func(i, j int) int64 {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p == 0 { // 首次计算
			*p = max(dfs(i-1, j), dfs(i-2, j^1)) + int64(c[j][i])
		}
		return *p
	}
	return max(dfs(n-1, 0), dfs(n-1, 1))
}

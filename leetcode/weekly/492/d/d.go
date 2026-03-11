package main

// https://space.bilibili.com/206214
func minCost(s string, encCost, flatCost int) int64 {
	n := len(s)
	sum := make([]int, n+1)
	for i, ch := range s {
		sum[i+1] = sum[i] + int(ch-'0')
	}

	// 计算子串 [l, r) 的最小费用，注意区间是左闭右开，方便计算
	var dfs func(int, int) int
	dfs = func(l, r int) int {
		x := sum[r] - sum[l]
		if x == 0 {
			return flatCost
		}

		// 不拆分
		res := (r - l) * x * encCost

		// 拆分
		if (r-l)%2 == 0 {
			m := (l + r) / 2
			res = min(res, dfs(l, m)+dfs(m, r))
		}

		return res
	}
	return int64(dfs(0, n))
}

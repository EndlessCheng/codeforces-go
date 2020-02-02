package copypasta

import "sort"

// https://oi-wiki.org/dp/

/*
若使用滚动数组，注意在下次复用时初始化第一排所有元素
但是实际情况是使用滚动数组仅降低了内存，执行效率与不使用时无异
*/

func dpCollections() {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	generalDP := func(a []int) (ans int) {
		n := len(a)
		dp := make([]int, n) // n+1
		for i := range dp {
			dp[i] = -1
		}
		var f func(int) int
		f = func(p int) (_ans int) {
			if p < 0 || p >= n {
				return 0
			}
			if dp[p] != -1 {
				return dp[p]
			}
			defer func() { dp[p] = _ans }()
			// 转移方程
			ap := a[p]
			_ = ap
			return
		}
		//return f(0)
		for i := range a {
			fi := f(i)
			ans = max(ans, fi)
		}
		return
	}

	generalDP2 := func(a []int) int {
		n := len(a)
		cost := func(l, r int) int {
			return 1
		}
		const mx int = 505
		dp := [mx][mx]int{}
		vis := [mx][mx]bool{}
		var f func(int, int) int
		f = func(l, r int) (_ans int) {
			if l >= r {
				return 0
			}
			if vis[l][r] {
				return dp[l][r]
			}
			vis[l][r] = true
			defer func() { dp[l][r] = _ans }()
			// 转移方程
			if a[l] == a[r] {
				return f(l+1, r-1)
			}
			f1 := f(l+1, r) + cost(l, r)
			f2 := f(l, r-1) + cost(l, r)
			return min(f1, f2)
		}
		return f(0, n-1)
	}

	// 由于数据范围的原因，采用 map 记忆化
	generalDPMap := func(x, y int) int {
		type pair struct{ x, y int }
		dp := map[pair]int{}
		var f func(int, int) int
		f = func(x, y int) (_ans int) {
			// 边界检查
			// ...
			p := pair{x, y}
			if v, ok := dp[p]; ok {
				return v
			}
			defer func() { dp[p] = _ans }()
			// 转移方程
			// ...
			return
		}
		return f(x, y)
	}

	// O(nlogn) LIS
	// https://oi-wiki.org/dp/basic/#_12
	lis := func(arr []int) int {
		dp := make([]int, 0, len(arr))
		for _, v := range arr {
			if i := sort.SearchInts(dp, v); i < len(dp) {
				dp[i] = v
			} else {
				dp = append(dp, v)
			}
		}
		return len(dp)
	}

	// 01背包
	// https://oi-wiki.org/dp/knapsack/
	knapsack01 := func(values, weights []int, maxW int) int {
		n := len(values)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, maxW+1)
		}
		for i, vi := range values {
			wi := weights[i]
			for j, dpij := range dp[i] {
				if j < wi {
					dp[i+1][j] = dpij
				} else {
					dp[i+1][j] = max(dpij, dp[i][j-wi]+vi)
				}
			}
		}
		return dp[n][maxW]
	}

	// TODO: 单调队列/单调栈优化
	// https://oi-wiki.org/dp/opt/monotonous-queue-stack/

	// TODO: 斜率优化
	// https://oi-wiki.org/dp/opt/slope/

	// TODO: 四边形不等式优化
	// https://oi-wiki.org/dp/opt/quadrangle/

	// 树上最大匹配
	// https://codeforces.com/blog/entry/2059
	// g[v] = ∑{max(f[son],g[son])}
	// f[v] = max{1+g[son]+g[v]−max(f[son],g[son])}
	maxMatchingOnTree := func(n int, g [][]int) int {
		cover, nonCover := make([]int, n), make([]int, n)
		var f func(int, int)
		f = func(v, fa int) {
			for _, w := range g[v] {
				if w != fa {
					f(w, v)
					nonCover[v] += max(cover[w], nonCover[w])
				}
			}
			for _, w := range g[v] {
				cover[v] = max(cover[v], 1+nonCover[w]+nonCover[v]-max(cover[w], nonCover[w]))
			}
		}
		f(0, -1)
		return max(cover[0], nonCover[0])
	}

	_ = []interface{}{
		generalDP, generalDP2, generalDPMap,
		lis, knapsack01,
		maxMatchingOnTree,
	}
}

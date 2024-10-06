package main

// https://space.bilibili.com/206214
func remainingMethods(n, k int, invocations [][]int) (ans []int) {
	g := make([][]int, n)
	for _, e := range invocations {
		g[e[0]] = append(g[e[0]], e[1])
	}

	// 标记所有可疑方法
	isSuspicious := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		isSuspicious[x] = true
		for _, y := range g[x] {
			if !isSuspicious[y] { // 避免无限递归
				dfs(y)
			}
		}
	}
	dfs(k)

	// 检查是否有【非可疑方法】->【可疑方法】的边
	for _, e := range invocations {
		if !isSuspicious[e[0]] && isSuspicious[e[1]] {
			// 无法移除可疑方法
			for i := range n {
				ans = append(ans, i)
			}
			return
		}
	}

	// 移除所有可疑方法
	for i, b := range isSuspicious {
		if !b {
			ans = append(ans, i)
		}
	}
	return
}

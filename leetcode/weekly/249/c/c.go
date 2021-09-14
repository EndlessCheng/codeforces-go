package main

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7

func colorTheGrid(m, n int) (ans int) {
	m *= 2
	// 预处理所有合法状态：用四进制表示颜色
	valid := []int{}
outer:
	for mask := 0; mask < 1<<m; mask++ {
		pre := 0
		for j := 0; j < m; j += 2 {
			color := mask >> j & 3
			if color == 0 || color == pre { // 未涂色或相邻颜色相同
				continue outer
			}
			pre = color
		}
		valid = append(valid, mask)
	}

	// 预处理所有合法状态能转移到哪些合法状态（记录合法状态的下标）
	to := make([][]int, len(valid))
	for i, v := range valid {
	o:
		for j, w := range valid {
			for k := 0; k < m; k += 2 {
				if v>>k&3 == w>>k&3 { // 相邻颜色相同
					continue o
				}
			}
			to[i] = append(to[i], j)
		}
	}

	// 滚动数组，用当前行更新下一行不同状态的方案数
	dp := make([]int, len(valid))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		tmp := dp
		dp = make([]int, len(valid))
		for j, dv := range tmp {
			for _, t := range to[j] {
				dp[t] = (dp[t] + dv) % mod
			}
		}
	}
	for _, dv := range dp {
		ans += dv
	}
	return ans % mod
}

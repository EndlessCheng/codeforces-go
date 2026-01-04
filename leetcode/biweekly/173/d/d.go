package main

// https://space.bilibili.com/206214
//func numberOfRoutes1(grid []string, d int) int {
//	const mod = 1_000_000_007
//	m := len(grid[0])
//	f := make([]int, m+1)
//	g := make([]int, m+1)
//
//	for i, row := range grid {
//		s := make([]int, m+1)
//		for j, ch := range row {
//			if ch == '#' {
//				s[j+1] = s[j]
//			} else if i == 0 { // 第一行（起点）
//				s[j+1] = s[j] + 1 // 原地不动，算一种方案
//			} else {
//				s[j+1] = (s[j] + f[min(j+d, m)] - f[max(j-d+1, 0)] + g[min(j+d, m)] - g[max(j-d+1, 0)]) % mod
//			}
//		}
//		f = s
//
//		for j, ch := range row {
//			if ch == '#' {
//				g[j+1] = g[j]
//			} else {
//				g[j+1] = (g[j] + f[min(j+d+1, m)] - f[j+1] + f[j] - f[max(j-d, 0)]) % mod
//			}
//		}
//	}
//
//	return (f[m] + g[m] + mod*2) % mod // +mod*2 保证结果非负
//}
func numberOfRoutes1(grid []string, d int) int {
	const mod = 1_000_000_007
	m := len(grid[0])
	sum := make([]int, m+1)

	for i, row := range grid {
		// 从 i-1 行移动到 i 行的方案数
		f := make([]int, m)
		for j, ch := range row {
			if ch == '#' {
				continue
			}
			if i == 0 { // 第一行（起点）
				f[j] = 1 // DP 初始值
			} else {
				f[j] = sum[min(j+d, m)] - sum[max(j-d+1, 0)]
			}
		}

		// f 的前缀和
		sumF := make([]int, m+1)
		for j, v := range f {
			sumF[j+1] = (sumF[j] + v) % mod
		}

		// 从 i 行移动到 i 行的方案数
		g := make([]int, m)
		for j, ch := range row {
			if ch == '#' {
				continue
			}
			// 不能原地不动，减去 f[j]
			g[j] = sumF[min(j+d+1, m)] - sumF[max(j-d, 0)] - f[j]
		}

		// f[j] + g[j] 的前缀和
		for j, fj := range f {
			sum[j+1] = (sum[j] + fj + g[j]) % mod
		}
	}

	return (sum[m] + mod) % mod // +mod 保证结果非负
}

func numberOfRoutes(grid []string, d int) int {
	const mod = 1_000_000_007
	m := len(grid[0])
	sumF := make([]int, m+1)
	sum := make([]int, m+1)

	for i, row := range grid {
		// f 的前缀和
		for j, ch := range row {
			if ch == '#' {
				sumF[j+1] = sumF[j]
			} else if i == 0 { // 第一行（起点）
				sumF[j+1] = sumF[j] + 1 // DP 初始值
			} else {
				sumF[j+1] = (sumF[j] + sum[min(j+d, m)] - sum[max(j-d+1, 0)]) % mod
			}
		}

		// f[j] + g[j] 的前缀和
		for j, ch := range row {
			if ch == '#' {
				sum[j+1] = sum[j]
			} else {
				// -f[j] 和 +f[j] 抵消了
				sum[j+1] = (sum[j] + sumF[min(j+d+1, m)] - sumF[max(j-d, 0)]) % mod
			}
		}
	}

	return (sum[m] + mod) % mod // +mod 保证结果非负
}

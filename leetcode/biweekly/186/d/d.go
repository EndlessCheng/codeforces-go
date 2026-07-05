package main

import "math"

// https://space.bilibili.com/206214
const mod = 1_000_000_007

// 115. 不同的子序列
func numDistinct(s, t string) int {
	n, m := len(s), len(t)
	if n < m {
		return 0
	}

	f := make([]int, m+1)
	f[0] = 1
	for i, x := range s {
		for j := min(i, m-1); j >= max(m-n+i, 0); j-- {
			if byte(x) == t[j] {
				f[j+1] = (f[j+1] + f[j]) % mod
			}
		}
	}
	return f[m]
}

func interleaveCharacters(word1, word2, target string) int {
	n, m1, m2 := len(target), len(word1), len(word2)
	f := make([][][]int, n+1)
	for i := range f {
		f[i] = make([][]int, m1+2)
		for j := range f[i] {
			f[i][j] = make([]int, m2+2)
		}
	}
	for j := 1; j < m1+2; j++ {
		for k := 1; k < m2+2; k++ {
			f[0][j][k] = 1
		}
	}

	for i, ch := range target {
		for j := range m1 + 1 {
			// j+k >= i+1
			for k := max(0, i+1-j); k <= m2; k++ {
				res := f[i+1][j][k+1] + f[i+1][j+1][k] - f[i+1][j][k]
				if j > 0 && word1[j-1] == byte(ch) {
					res += f[i][j][k+1] - f[i][j][k]
				}
				if k > 0 && word2[k-1] == byte(ch) {
					res += f[i][j+1][k] - f[i][j][k]
				}
				f[i+1][j+1][k+1] = res % mod
			}
		}
	}

	ans := f[n][m1+1][m2+1] - numDistinct(word1, target) - numDistinct(word2, target)
	return (ans%mod + mod) % mod // 保证 ans 非负
}

func interleaveCharacters1(word1, word2, target string) int {
	n, m1, m2 := len(target), len(word1), len(word2)

	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, m1+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, m2+1)
			for p := range memo[i][j] {
				memo[i][j][p] = math.MinInt
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if j < -1 || k < -1 || j+k+2 < i+1 {
			return 0
		}
		if i < 0 {
			return 1
		}
		p := &memo[i][j+1][k+1]
		if *p != math.MinInt {
			return *p
		}

		// 不选 word1[j] 或 word2[k]（至少一个不在 target 中）
		res := dfs(i, j-1, k) + dfs(i, j, k-1) - dfs(i, j-1, k-1) // 容斥

		// 选 word1[j] 和 word2[k]（都在 target 中）
		if j >= 0 && word1[j] == target[i] {
			// 选 word1[j]，减去不选 word2[k] 的方案数，就是 word2[k] 也在 target 中的方案数
			res += dfs(i-1, j-1, k) - dfs(i-1, j-1, k-1)
		}
		if k >= 0 && word2[k] == target[i] {
			// 选 word2[k]，减去不选 word1[j] 的方案数，就是 word1[j] 也在 target 中的方案数
			res += dfs(i-1, j, k-1) - dfs(i-1, j-1, k-1)
		}

		res %= mod
		*p = res
		return res
	}

	ans := dfs(n-1, m1-1, m2-1) - numDistinct(word1, target) - numDistinct(word2, target)
	return (ans%mod + mod) % mod // 保证 ans 非负
}

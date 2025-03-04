package main

// https://space.bilibili.com/206214
func longestPalindromicSubsequence(s string, K int) int {
	n := len(s)
	cnt := 0
	for i := range n / 2 {
		d := abs(int(s[i]) - int(s[n-1-i]))
		cnt += min(d, 26-d)
	}
	if cnt <= K {
		return n
	}

	f := make([][][]int, K+1)
	for k := range f {
		f[k] = make([][]int, n)
		for j := range f[k] {
			f[k][j] = make([]int, n)
		}
		for i := n - 1; i >= 0; i-- {
			f[k][i][i] = 1
			for j := i + 1; j < n; j++ {
				res := max(f[k][i+1][j], f[k][i][j-1])
				d := abs(int(s[i]) - int(s[j]))
				op := min(d, 26-d)
				if op <= k {
					res = max(res, f[k-op][i+1][j-1]+2)
				}
				f[k][i][j] = res
			}
		}
	}
	return f[K][0][n-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func longestPalindromicSubsequence1(s string, k int) int {
	n := len(s)
	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, k+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i >= j {
			return j - i + 1 // i=j+1 时返回 0，i=j 时返回 1
		}
		p := &memo[i][j][k]
		if *p != -1 {
			return *p
		}
		res := max(dfs(i+1, j, k), dfs(i, j-1, k))
		d := abs(int(s[i]) - int(s[j]))
		op := min(d, 26-d)
		if op <= k {
			res = max(res, dfs(i+1, j-1, k-op)+2)
		}
		*p = res
		return res
	}
	return dfs(0, n-1, k)
}

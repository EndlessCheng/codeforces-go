package main

// https://space.bilibili.com/206214
func largestPalindrome(n, k int) string {
	pow10 := make([]int, n)
	pow10[0] = 1
	for i := 1; i < n; i++ {
		pow10[i] = pow10[i-1] * 10 % k
	}

	ans := make([]byte, n)
	m := (n + 1) / 2
	vis := make([][]bool, m+1)
	for i := range vis {
		vis[i] = make([]bool, k)
	}
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i == m {
			return j == 0
		}
		vis[i][j] = true
		for d := 9; d >= 0; d-- { // 贪心：从大到小枚举
			var j2 int
			if n%2 > 0 && i == m-1 { // 正中间
				j2 = (j + d*pow10[i]) % k
			} else {
				j2 = (j + d*(pow10[i]+pow10[n-1-i])) % k
			}
			if !vis[i+1][j2] && dfs(i+1, j2) {
				ans[i] = '0' + byte(d)
				ans[n-1-i] = ans[i]
				return true
			}
		}
		return false
	}
	dfs(0, 0)
	return string(ans)
}

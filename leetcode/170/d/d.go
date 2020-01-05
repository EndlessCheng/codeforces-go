package main

func minInsertions(ss string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	s := []byte(ss)
	const mx = 505
	dp := [mx][mx]int{}
	vis := [mx][mx]bool{}
	var f func(l, r int) int
	f = func(l, r int) int {
		if l >= r {
			return 0
		}
		if vis[l][r] {
			return dp[l][r]
		}
		vis[l][r] = true
		var ans int
		if s[l] == s[r] {
			ans = f(l+1, r-1)
		} else {
			ans = min(f(l+1, r)+1, f(l, r-1)+1)
		}
		dp[l][r]=ans
		return ans
	}
	return f(0, len(s)-1)
}

package main

// github.com/EndlessCheng/codeforces-go
func beautySum(s string) (ans int) {
	n := len(s)
	sum := make([][26]int, n+1)
	for i, b := range s {
		sum[i+1] = sum[i]
		sum[i+1][b-'a']++
	}
	for l := range s {
		for r := l + 1; r <= n; r++ {
			mi, mx := n, 0
			for i := 0; i < 26; i++ {
				if d := sum[r][i] - sum[l][i]; d > 0 {
					mi = min(mi, d)
					mx = max(mx, d)
				}
			}
			ans += mx - mi
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

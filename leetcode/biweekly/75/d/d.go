package main

// github.com/EndlessCheng/codeforces-go
func sumScores(s string) int64 {
	n := len(s)
	ans := n
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		z[i] = max(0, min(z[i-l], r-i+1))
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
		ans += z[i]
	}
	return int64(ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }

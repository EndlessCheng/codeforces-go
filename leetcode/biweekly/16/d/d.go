package main

// github.com/EndlessCheng/codeforces-go
func pathsWithMaxScore(a []string) (ans []int) {
	type pair struct {
		s, c int
		ok   bool
	}
	const mod int = 1e9 + 7
	n, m := len(a), len(a[0])
	dp := make([][]pair, n+1)
	for i := range dp {
		dp[i] = make([]pair, m+1)
	}
	dp[0][0] = pair{0, 1, true}
	for i, r := range a {
		if i == 0 {
			for j := 1; j < m; j++ {
				if r[j] != 'X' {
					dp[i][j] = dp[i][j-1]
					dp[i][j].s += int(r[j] & 15)
				}
			}
			continue
		}
		for j, b := range r {
			if b == 'X' {
				continue
			}
			if j == 0 {
				dp[i][j] = dp[i-1][j]
				dp[i][j].s += int(r[j] & 15)
				continue
			}
			p := dp[i-1][j-1]
			for _, q := range []pair{dp[i-1][j], dp[i][j-1]} {
				if q.s > p.s {
					p = q
				} else if q.s == p.s {
					p.c += q.c
				}
				p.ok = p.ok || q.ok
			}
			if i < n-1 || j < m-1 {
				p.s += int(b & 15)
			}
			p.c %= mod
			dp[i][j] = p
		}
	}
	p := dp[n-1][m-1]
	if !p.ok { // 由于取模的原因，用 p.c == 0 来判断是错误的
		p.s = 0
	}
	return []int{p.s, p.c}
}

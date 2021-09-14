package main

// github.com/EndlessCheng/codeforces-go
func minCost(s string, cost []int) (ans int) {
	for i, n := 0, len(s); i < n; {
		v, max := s[i], 0
		for ; i < n && s[i] == v; i++ {
			w := cost[i]
			ans += w
			if w > max {
				max = w
			}
		}
		ans -= max
	}
	return
}

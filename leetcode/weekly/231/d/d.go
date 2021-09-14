package main

// github.com/EndlessCheng/codeforces-go
func minChanges(a []int, k int) (ans int) {
	cnt := make([][1024]int, k)
	n := len(a)
	mi := n
	for i := 0; i < k; i++ {
		mx := 0
		for j := i; j < n; j += k {
			cnt[i][a[j]]++
			mx = max(mx, cnt[i][a[j]])
		}
		ans += mx
		mi = min(mi, mx)
	}
	ans -= mi

	dp := make([][1024]int, k-1)
	for i := range dp {
		for j := 0; j < 1024; j++ {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(p, xor int) (res int) {
		if p == k-1 {
			return cnt[p][xor]
		}
		dv := &dp[p][xor]
		if *dv >= 0 {
			return *dv
		}
		defer func() { *dv = res }()
		for i := p; i < n; i += k {
			res = max(res, cnt[p][a[i]]+f(p+1, xor^a[i]))
		}
		return
	}
	return n - max(ans, f(0, 0))
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

package main

// github.com/EndlessCheng/codeforces-go
func minMoves(a []int, lim int) (ans int) {
	n := len(a)
	d := make([]int, 2*lim+2)
	cnt := make([]int, 2*lim+1)
	for i, v := range a[:n/2] {
		w := a[n-1-i]
		l, r := 1+min(v, w), lim+max(v, w)
		d[l]++
		d[v+w]--
		d[v+w+1]++
		d[r+1]--
		cnt[v+w]++
	}
	ans = 1e9
	one := 0
	for s := 2; s <= 2*lim; s++ {
		one += d[s]
		ans = min(ans, n-one-2*cnt[s])
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

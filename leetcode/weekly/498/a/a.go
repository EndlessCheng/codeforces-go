package main

// https://space.bilibili.com/206214
func firstStableIndex(a []int, k int) (ans int) {
	n := len(a)
	suf := make([]int, n)
	suf[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = min(suf[i+1], a[i])
	}
	pre := 0
	for i, v := range a {
		pre = max(pre, v)
		if pre - suf[i] <= k {
			return i
		}
	}
	return -1
}

package main

// https://space.bilibili.com/206214
func longestSubarray(a []int) (ans int) {
	mx := int(-1e9)
	for _, v := range a {
		mx = max(mx, v)
	}
	for i, n := 0, len(a); i < n; {
		st := i
		if a[i] != mx {
			i++
			continue
		}
		for ; i < n && a[i] == mx; i++ {

		}
		ans = max(ans, i-st)
	}
	return
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

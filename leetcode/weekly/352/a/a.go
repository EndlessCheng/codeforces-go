package main

// https://space.bilibili.com/206214
func longestAlternatingSubarray(a []int, threshold int) (ans int) {
	for i, n := 0, len(a); i < n; {
		if a[i]%2 > 0 || a[i] > threshold {
			i++
		} else {
			i0 := i
			for i++; i < n && a[i] <= threshold && a[i]%2 != a[i-1]%2; i++ {}
			ans = max(ans, i-i0)
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }

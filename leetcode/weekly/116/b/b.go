package main

// https://space.bilibili.com/206214
func maxWidthRamp(a []int) (ans int) {
	n := len(a)
	st := []int{0}
	for i := 1; i < n; i++ {
		if a[i] < a[st[len(st)-1]] {
			st = append(st, i)
		}
	}
	for i := n-1; i > 0; i-- {
		for len(st) > 0 && a[i] >= a[st[len(st)-1]] {
			ans = max(ans, i-st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

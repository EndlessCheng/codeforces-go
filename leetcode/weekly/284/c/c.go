package main

// github.com/EndlessCheng/codeforces-go
func maximumTop(a []int, k int) (ans int) {
	n := len(a)
	if n == 1 || k == 0 {
		if k%2 == 1 { return -1 }
		return a[0]
	}
	// 删除 a[k-1] 以及 a[k+1:]，下面直接取 a 的最大值
	if k < n {
		a = append(a[:k-1], a[k])
	} else if k == n {
		a = a[:n-1]
	}
	for _, v := range a {
		if v > ans { ans = v }
	}
	return
}

package main

// github.com/EndlessCheng/codeforces-go
func waysToMakeFair(a []int) (ans int) {
	const k = 2
	n := len(a)
	for len(a)%k > 0 {
		a = append(a, 0)
	}
	sum := make([]int, len(a)+k)
	for i, v := range a {
		sum[i+k] = sum[i] + v
	}
	pre := func(x, m int) int {
		if x%k <= m {
			return sum[x/k*k+m]
		}
		return sum[(x+k-1)/k*k+m]
	}
	query := func(l, r, m int) int {
		return pre(r, m) - pre(l, m)
	}
	// 由于 a 发生了变化，这里用初始长度
	for i := 0; i < n; i++ {
		if query(0, i, 0)+query(i+1, n, 1) == query(0, i, 1)+query(i+1, n, 0) {
			ans++
		}
	}
	return
}

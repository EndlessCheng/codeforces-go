package main

// github.com/EndlessCheng/codeforces-go
func oddnumber(n, m int, x, y []int) (ans int) {
	d := make([]int, n+2)
	for i, l := range x {
		d[l]++
		d[y[i]+1]--
	}
	v := m
	for i := 1; i <= n; i++ {
		v += d[i]
		ans += v & 1
	}
	return
}

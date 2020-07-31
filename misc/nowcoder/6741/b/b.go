package main

// github.com/EndlessCheng/codeforces-go
func tree5(a []int) int64 {
	n, s, i := len(a), 0, 0
	var f func(int) int
	f = func(p int) int {
		v := a[i]
		i++
		if p*2 <= n {
			s += v ^ f(p*2)
		}
		if p*2+1 <= n {
			s += v ^ f(p*2+1)
		}
		return v
	}
	f(1)
	return int64(s)
}

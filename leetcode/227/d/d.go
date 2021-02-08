package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minAbsDifference(a []int, g int) (ans int) {
	ans = abs(g)

	n := len(a)
	if n == 1 {
		return min(ans, abs(a[0]-g))
	}

	x := subSum(a[:n/2])

	k := sort.SearchInts(x, g)
	if k < len(x) {
		ans = min(ans, abs(x[k]-g))
	}
	if k > 0 {
		ans = min(ans, abs(x[k-1]-g))
	}

	b := a[n/2:]
	y := make([]int, 1<<len(b))
	for i, v := range b {
		for j := 0; j < 1<<i; j++ {
			y[1<<i|j] = y[j] + v
			s := y[1<<i|j] - g
			k := sort.SearchInts(x, -s)
			if k < len(x) {
				ans = min(ans, abs(s+x[k]))
			}
			if k > 0 {
				ans = min(ans, abs(s+x[k-1]))
			}
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(a, b []int) []int {
	i, n := 0, len(a)
	j, m := 0, len(b)
	res := make([]int, 0, n+m)
	for {
		if i == n {
			return append(res, b[j:]...)
		}
		if j == m {
			return append(res, a[i:]...)
		}
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
}

func subSum(a []int) []int {
	x := []int{0}
	for _, v := range a {
		b := make([]int, len(x))
		for i, w := range x {
			b[i] = w + v
		}
		x = merge(x, b)
	}
	return x
}

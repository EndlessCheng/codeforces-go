package main

// github.com/EndlessCheng/codeforces-go
func findLatestStep(a []int, m int) (ans int) {
	n := len(a)
	fa := make([]int, n+1)
	sz := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	cnt := make([]int, n+1)
	merge := func(from, to int) {
		from, to = f(from), f(to)
		if from != to {
			cnt[sz[from]]--
			cnt[sz[to]]--
			sz[to] += sz[from]
			cnt[sz[to]]++
			fa[from] = to
		}
	}
	ans = -1
	for i, v := range a {
		merge(v, v+1)
		if cnt[m+1] > 0 {
			ans = i + 1
		}
	}
	return
}

package main

// github.com/EndlessCheng/codeforces-go
func smallestMissingValueSubtree(parents []int, a []int) []int {
	n := len(parents)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		v := parents[w]
		g[v] = append(g[v], w)
	}
	mex := make([]int, n)
	var f func(int) map[int]bool
	f = func(v int) map[int]bool {
		set := map[int]bool{}
		mex[v] = 1
		for _, w := range g[v] {
			s := f(w)
			if len(s) > len(set) {
				set, s = s, set
			}
			for x := range s {
				set[x] = true
			}
			if mex[w] > mex[v] {
				mex[v] = mex[w]
			}
		}
		set[a[v]] = true
		for set[mex[v]] {
			mex[v]++
		}
		return set
	}
	f(0)
	return mex
}

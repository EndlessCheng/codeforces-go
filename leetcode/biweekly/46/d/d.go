package main

// github.com/EndlessCheng/codeforces-go
type pair struct{ d, v int }

func getCoprimes(a []int, edges [][]int) (ans []int) {
	n := len(a)
	ans = make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	vs := [51][]pair{}
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		val := a[v]
		mxD := 0
		for i, ps := range vs[:] {
			if len(ps) > 0 && gcd(i, val) == 1 && ps[len(ps)-1].d > mxD {
				mxD = ps[len(ps)-1].d
				ans[v] = ps[len(ps)-1].v
			}
		}
		vs[val] = append(vs[val], pair{d, v})
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
		vs[val] = vs[val][:len(vs[val])-1]
	}
	f(0, -1, 1)
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

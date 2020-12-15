package main

// github.com/EndlessCheng/codeforces-go
func MaxDiameter(n int, vs, ws []int) (ans int) {
	g := make([][]int, n)
	deg := make([]int, n)
	for i, v := range vs {
		w := ws[i]
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		deg[v]++
		deg[w]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 1 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 1 {
				q = append(q, w)
			}
		}
	}

	f := func(root int) int {
		maxD, u := -1, 0
		var f func(v, fa, d int)
		f = func(v, fa, d int) {
			if d > maxD {
				maxD, u = d, v
			}
			for _, w := range g[v] {
				if w != fa && deg[w] < 2 {
					f(w, v, d+1)
				}
			}
		}
		f(root, -1, 0)
		maxDep := maxD
		ans = max(ans, maxDep)
		maxD = -1
		f(u, -1, 0)
		ans = max(ans, maxD)
		return maxDep
	}
	ds := []int{}
	for root, d := range deg {
		if d > 1 {
			maxD := f(root)
			ds = append(ds, maxD)
			ans = max(ans, maxD)
		}
	}

	n = len(ds)
	ds = append(ds, ds...)
	for i := 0; i < n; i++ {
		for j := i + 1; j < i+n; j++ {
			ans = max(ans, ds[i]+j-i+ds[j])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

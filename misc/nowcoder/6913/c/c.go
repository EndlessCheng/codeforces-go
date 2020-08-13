package main

// github.com/EndlessCheng/codeforces-go
func solve(n, q int, a, vs, ws, x, y []int) []int {
	const mod int = 1e9 + 7
	g := make([][]int, n)
	for i, v := range vs {
		v--
		w := ws[i] - 1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	pa := make([]int, n)
	for i := range pa {
		pa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if pa[x] != x {
			pa[x] = find(pa[x])
		}
		return pa[x]
	}
	pow := func(x int) int {
		res := 1
		for n := mod - 2; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	div := func(a, b int) int { return a * pow(b) % mod }

	ans := make([]int, q)
	type query struct{ w, i int }
	qs := make([][]query, n)
	for i, v := range x {
		v--
		w := y[i] - 1
		if v != w {
			qs[v] = append(qs[v], query{w, i})
			qs[w] = append(qs[w], query{v, i})
		} else {
			ans[i] = a[v]
		}
	}

	mul := make([]int, n)
	zero := make([]int, n)
	vis := make([]int8, n)
	var f func(v, m, z int)
	f = func(v, m, z int) {
		mul[v] = m
		zero[v] = z
		vis[v] = 1
		for _, w := range g[v] {
			if vis[w] == 0 {
				if a[w] == 0 {
					f(w, m, z+1)
				} else {
					f(w, m*a[w]%mod, z)
				}
				pa[w] = v
			}
		}
		for _, q := range qs[v] {
			if w := q.w; vis[w] == 2 {
				lca := find(w)
				z := zero[v] + zero[w] - zero[lca]<<1
				if a[lca] == 0 {
					z++
				}
				if z == 0 {
					ans[q.i] = div(mul[v]*mul[w]%mod, mul[lca]*mul[lca]%mod) * a[lca] % mod
					// 如果 a[i] 可以是负数的话最后还要补上 +mod%mod
				}
			}
		}
		vis[v] = 2
	}
	if a[0] == 0 {
		f(0, 1, 1)
	} else {
		f(0, a[0], 0)
	}
	return ans
}

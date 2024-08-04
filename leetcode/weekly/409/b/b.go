package main

// https://space.bilibili.com/206214
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	from := make([][]int, n)
	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = i
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		from[r] = append(from[r], l)
		if f[l]+1 < f[r] {
			f[r] = f[l] + 1
			for i := r + 1; i < n; i++ {
				f[i] = min(f[i], f[i-1]+1)
				for _, j := range from[i] {
					f[i] = min(f[i], f[j]+1)
				}
			}
		}
		ans[qi] = f[n-1]
	}
	return ans
}

func shortestDistanceAfterQueries2(n int, queries [][]int) []int {
	g := make([][]int, n-1)
	for i := range g {
		g[i] = append(g[i], i+1)
	}

	vis := make([]int, n-1)
	bfs := func(i int) int {
		q := []int{0}
		for step := 1; ; step++ {
			tmp := q
			q = nil
			for _, x := range tmp {
				for _, y := range g[x] {
					if y == n-1 {
						return step
					}
					if vis[y] != i {
						vis[y] = i
						q = append(q, y)
					}
				}
			}
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		g[q[0]] = append(g[q[0]], q[1])
		ans[i] = bfs(i + 1)
	}
	return ans
}

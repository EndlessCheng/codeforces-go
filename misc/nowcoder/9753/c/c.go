package main

// github.com/EndlessCheng/codeforces-go
func PointsOnDiameter(n int, u, v []int) (ans int) {
	g := make([][]int, n+1)
	for i, x := range u {
		y := v[i]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	L := 0
	var f func(v, p int) (int, int)
	f = func(v, p int) (int, int) {
		mxD, cnt := 0, 0
		for _, w := range g[v] {
			if w != p {
				d, c := f(w, v)
				if l := mxD + d; l > L {
					L, ans = l, cnt+c+1
				} else if l == L {
					ans += c
				}
				if d > mxD {
					mxD, cnt = d, c
				} else if d == mxD {
					cnt += c
				}
			}
		}
		return mxD + 1, cnt + 1
	}
	f(1, 0)
	return
}

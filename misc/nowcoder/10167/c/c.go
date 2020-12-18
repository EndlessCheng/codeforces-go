package main

// github.com/EndlessCheng/codeforces-go
func solve(n int, vs, ws, a []int) (ans int64) {
	g := make([][]int, n)
	for i, v := range vs {
		w := ws[i]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	for i := 0; i < 21; i++ {
		cnt := 0
		var f func(int, int) int
		f = func(v, fa int) int {
			one := a[v] >> i & 1
			cnt += one
			for _, w := range g[v] {
				if w != fa {
					o := f(w, v)
					if one > 0 {
						cnt += one * o
						one += o
					}
				}
			}
			return one
		}
		f(0, n)
		ans += 1 << i * int64(cnt)
	}
	return
}

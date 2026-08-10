package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2244F(in io.Reader, out io.Writer) {
	var T, n, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for w := 1; w < n; w++ {
			Fscan(in, &p)
			p--
			g[p] = append(g[p], w)
		}
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ok := true
		type pair struct{ l, r int }
		var dfs func(int) pair
		dfs = func(v int) pair {
			if g[v] == nil {
				return pair{a[v], a[v]}
			}
			ps := []pair{}
			for _, w := range g[v] {
				ps = append(ps, dfs(w))
			}
			idx := -1
			m := len(ps)
			if ps[m-1].r >= ps[0].l {
				idx = 0
			}
			for i := 1; i < m; i++ {
				if ps[i-1].r > ps[i].l {
					if idx >= 0 {
						ok = false
					}
					idx = i
				}
			}
			return pair{ps[idx].l, ps[(idx-1+m)%m].r}
		}
		dfs(0)

		if ok {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf2244F(bufio.NewReader(os.Stdin), os.Stdout) }

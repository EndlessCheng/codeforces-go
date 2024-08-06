package main

import (
	. "fmt"
	"io"
)

func cf1822F(in io.Reader, out io.Writer) {
	var T, n, k, c, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &c)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		nodes := make([]struct{ maxD, maxD2, my int }, n)
		var dfs func(int, int) int
		dfs = func(x, fa int) int {
			p := &nodes[x]
			for _, y := range g[x] {
				if y == fa {
					continue
				}
				d := dfs(y, x) + k
				if d > p.maxD {
					p.maxD2 = p.maxD
					p.maxD = d
					p.my = y
				} else if d > p.maxD2 {
					p.maxD2 = d
				}
			}
			return p.maxD
		}
		ans := dfs(0, -1)

		var reroot func(int, int, int, int)
		reroot = func(x, fa, fromUp, cost int) {
			p := nodes[x]
			ans = max(ans, max(fromUp, p.maxD)-cost)
			for _, y := range g[x] {
				if y == fa {
					continue
				}
				exceptY := p.maxD
				if y == p.my {
					exceptY = p.maxD2
				}
				reroot(y, x, max(fromUp, exceptY)+k, cost+c)
			}
		}
		reroot(0, -1, 0, 0)
		Fprintln(out, ans)
	}
}

//func main() { cf1822F(bufio.NewReader(os.Stdin), os.Stdout) }

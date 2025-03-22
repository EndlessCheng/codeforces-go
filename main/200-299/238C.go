package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf238C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type nb struct{ to, inv int }
	g := make([][]nb, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], nb{w, 0})
		g[w] = append(g[w], nb{v, 1})
	}

	ans := n
	for i := range n {
		inv := 0
		var dfs func(int, int) int
		dfs = func(v, fa int) (res int) {
			for _, e := range g[v] {
				w := e.to
				if w != fa {
					inv += e.inv
					res = max(res, dfs(w, v)+e.inv)
				}
			}
			return
		}
		res := dfs(i, -1)
		ans = min(ans, inv-res)
	}
	Fprint(out, ans)
}

//func main() { cf238C(bufio.NewReader(os.Stdin), os.Stdout) }

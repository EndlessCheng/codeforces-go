package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1088E(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	mx := int(-1e18)
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		res := a[v]
		for _, w := range g[v] {
			if w != fa {
				res += dfs(w, v)
			}
		}
		mx = max(mx, res)
		return max(res, 0)
	}
	dfs(0, -1)

	dfs = func(v, fa int) int {
		res := a[v]
		for _, w := range g[v] {
			if w != fa {
				res += dfs(w, v)
			}
		}
		if res == mx {
			ans++
			return 0
		}
		return max(res, 0)
	}
	dfs(0, -1)

	Fprint(out, mx*ans, ans)
}

//func main() { cf1088E(bufio.NewReader(os.Stdin), os.Stdout) }

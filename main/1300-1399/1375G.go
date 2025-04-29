package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1375G(in io.Reader, out io.Writer) {
	var n, cnt int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	var dfs func(int, int, int)
	dfs = func(v, fa, d int) {
		cnt += d
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, d^1)
			}
		}
	}
	dfs(1, 0, 0)
	Fprint(out, min(cnt, n-cnt)-1)
}

//func main() { cf1375G(bufio.NewReader(os.Stdin), os.Stdout) }

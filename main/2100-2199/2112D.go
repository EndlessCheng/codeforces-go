package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2112D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n+1)
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		rt := 0
		for i, es := range g {
			if len(es) == 2 {
				rt = i
				break
			}
		}
		if rt == 0 {
			Fprintln(out, "NO")
			continue
		}

		Fprintln(out, "YES")
		var dfs func(int, int, bool)
		dfs = func(v, fa int, rev bool) {
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				if rev {
					Fprintln(out, w, v)
				} else {
					Fprintln(out, v, w)
				}
				dfs(w, v, !rev)
			}
		}
		dfs(rt, g[rt][1], false)
		dfs(rt, g[rt][0], true)
	}
}

//func main() { cf2112D(bufio.NewReader(os.Stdin), os.Stdout) }

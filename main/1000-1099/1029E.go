package main

import (
	. "fmt"
	"io"
)

func cf1029E(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		has1, covered := false, false
		for _, w := range g[v] {
			if w != fa {
				r := dfs(w, v)
				if r == 1 {
					has1 = true
				} else if r < 0 {
					covered = true
				}
			}
		}
		if v == 0 || fa == 0 {
			return 0
		}
		if has1 {
			ans++
			return -1
		}
		if covered {
			return 0
		}
		return 1
	}
	dfs(0, -1)
	Fprint(out, ans)
}

//func main() { cf1029E(bufio.NewReader(os.Stdin), os.Stdout) }

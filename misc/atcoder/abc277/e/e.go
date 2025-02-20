package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m, &k)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}
	sp := make([]bool, n)
	for ; k > 0; k-- {
		var v int
		Fscan(in, &v)
		sp[v-1] = true
	}

	vis := make([][2]bool, n)
	vis[0][0] = true
	type pair struct{ v, c int }
	q := []pair{{}}
	for step := 0; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			if p.v == n-1 {
				Fprint(out, step)
				return
			}
			v, c := p.v, p.c
			for _, e := range g[v] {
				w := e.to
				if e.wt != c && !vis[w][c] {
					vis[w][c] = true
					q = append(q, pair{w, c})
				}
				if sp[v] && e.wt == c && !vis[w][c^1] {
					vis[w][c^1] = true
					q = append(q, pair{w, c ^ 1})
				}
			}
		}
	}
	Fprint(out, -1)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, maxL, v, w, wt int
	Fscan(in, &n, &m, &maxL)
	type to struct{ to, wt int }
	g := make([][]to, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		g[v-1] = append(g[v-1], to{w - 1, wt})
	}

	fa := make([]int, n)
	f := func(maxW int) bool {
		for i := range fa {
			fa[i] = -1
		}
		vis := make([]bool, len(g))
		vis[0] = true
		type pair struct{ v, d int }
		q := []pair{{}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			v, d := p.v, p.d
			if v == n-1 {
				return d <= maxL
			}
			for _, e := range g[v] {
				if w := e.to; e.wt <= maxW && !vis[w] {
					vis[w] = true
					fa[w] = v
					q = append(q, pair{w, d + 1})
				}
			}
		}
		return false
	}
	maxW := sort.Search(1e9+1, f)
	if maxW > 1e9 {
		Fprint(out, -1)
		return
	}
	f(maxW)
	path := []int{}
	for v := n - 1; v != -1; v = fa[v] {
		path = append(path, v+1)
	}
	Fprintln(out, len(path)-1)
	for i := len(path) - 1; i >= 0; i-- {
		Fprint(out, path[i], " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }

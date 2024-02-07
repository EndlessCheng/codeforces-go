package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1927F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v, w, wt int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		type nb struct{ to, wt, rid int }
		g := make([][]nb, n)
		deg := make([]int, n)
		for i := 0; i < m; i++ {
			Fscan(in, &v, &w, &wt)
			v--
			w--
			g[v] = append(g[v], nb{w, wt, len(g[w])})
			g[w] = append(g[w], nb{v, wt, len(g[v]) - 1})
			deg[v]++
			deg[w]++
		}
		dfn := make([]int, n)
		dfsClock := 0
		var tarjan func(int, int) int
		tarjan = func(v, fa int) int {
			dfsClock++
			dfn[v] = dfsClock
			lowV := dfsClock
			for i, e := range g[v] {
				if w := e.to; dfn[w] == 0 {
					lowW := tarjan(w, v)
					lowV = min(lowV, lowW)
					if lowW > dfn[v] {
						g[v][i].wt = 1e9
						g[w][e.rid].wt = 1e9
					}
				} else if w != fa {
					lowV = min(lowV, dfn[w])
				}
			}
			return lowV
		}
		for v, timestamp := range dfn {
			if timestamp == 0 {
				tarjan(v, -1)
			}
		}

		mnV := 0
		mnWt := int(1e9)
		for v, r := range g {
			if r == nil {
				continue
			}
			sort.Slice(r, func(i, j int) bool { return r[i].wt < r[j].wt })
			if r[0].wt < mnWt {
				mnWt = r[0].wt
				mnV = v
			}
		}

		vis := make([]bool, n)
		st := []any{}
		var f func(int, int) bool
		f = func(v, fa int) bool {
			vis[v] = true
			st = append(st, v+1)
			for _, e := range g[v] {
				w := e.to
				if w != fa && (w == mnV || !vis[w] && f(w, v)) {
					return true
				}
			}
			st = st[:len(st)-1]
			return false
		}
		f(mnV, -1)
		Fprintln(out, mnWt, len(st))
		Fprintln(out, st...)
	}
}

//func main() { cf1927F(os.Stdin, os.Stdout) }

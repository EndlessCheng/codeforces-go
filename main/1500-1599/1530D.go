package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1530D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, st int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([]int, n)
		deg := make([]int, n)
		for i := range g {
			Fscan(in, &g[i])
			g[i]--
			deg[g[i]]++
		}

		rg := make([][]int, n)
		f := make([]int, n)
		from := make([]int, n)
		for i := range from {
			from[i] = -1
		}
		q := []int{}
		for i, d := range deg {
			if d == 0 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			w := g[v]
			rg[w] = append(rg[w], v)
			f[v]++
			if f[v] > f[w] {
				f[w] = f[v]
				from[w] = v
			}
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}

		type path struct{ st, end int }
		chain := []path{}
		var rdfs func(int, bool)
		rdfs = func(v int, change bool) {
			if rg[v] == nil {
				st = v
			} else {
				for _, w := range rg[v] {
					if w != from[v] {
						rdfs(w, true)
					}
				}
				rdfs(from[v], false)
			}
			if change {
				chain = append(chain, path{st, v})
			}
		}
		for i0, d := range deg {
			if d == 0 {
				continue
			}
			ring := []int{i0}
			for v := g[i0]; v != i0; v = g[v] {
				deg[v] = 0
				ring = append(ring, v)
			}
			for i, v := range ring {
				if rg[v] != nil {
					rdfs(v, false)
					chain = append(chain, path{st, ring[(i-1+len(ring))%len(ring)]})
				}
			}
		}

		for i, p := range chain {
			g[p.end] = chain[(i+1)%len(chain)].st
		}

		Fprintln(out, n-len(chain))
		for _, v := range g {
			Fprint(out, v+1, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1530D(os.Stdin, os.Stdout) }

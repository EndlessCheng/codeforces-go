package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF243B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type edge struct{ v, w int }

	var n, m, h, t, v, w int
	Fscan(in, &n, &m, &h, &t)
	g := make([][]int, n+1)
	edges := make([]edge, m)
	for i := range edges {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		edges[i] = edge{v, w}
	}
	for i := range g {
		sort.Ints(g[i])
	}
	for _, e := range edges {
		v, w := e.v, e.w
		vs, ws := g[v], g[w]
		if h < t && len(vs) > len(ws) || h > t && len(vs) < len(ws) {
			vs, ws = ws, vs
			v, w = w, v
		}
		if len(vs)-1 < h || len(ws)-1 < t {
			continue
		}
		var hs, ts, same []interface{}
		i, n := 0, len(vs)
		j, m := 0, len(ws)
		for i < n || j < m {
			if i == n {
				if y := ws[j]; y != v {
					ts = append(ts, y)
				}
				j++
			} else if j == m {
				if x := vs[i]; x != w {
					hs = append(hs, x)
				}
				i++
			} else {
				x, y := vs[i], ws[j]
				if x == w {
					i++
					continue
				}
				if y == v {
					j++
					continue
				}
				if x < y {
					if len(hs) < h {
						hs = append(hs, x)
					}
					i++
				} else if x > y {
					if len(ts) < t {
						ts = append(ts, y)
					}
					j++
				} else {
					same = append(same, x)
					i++
					j++
				}
			}
			if len(hs)+len(ts)+len(same) >= h+t {
				Fprintln(out, "YES")
				Fprintln(out, v, w)
				if len(hs) < h {
					popL := same[:h-len(hs)]
					same = same[h-len(hs):]
					hs = append(hs, popL...)
				}
				Fprintln(out, hs...)
				if len(ts) < t {
					ts = append(ts, same[:t-len(ts)]...)
				}
				Fprintln(out, ts...)
				return
			}
		}
	}
	Fprint(out, "NO")
}

//func main() { CF243B(os.Stdin, os.Stdout) }

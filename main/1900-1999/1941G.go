package main

import (
	. "fmt"
	"io"
)

func cf1941G(in io.Reader, out io.Writer) {
	var T, n, m, st, end int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		type pair struct{ v, c int }
		g := map[pair][]pair{}
		add := func(v, c int) {
			p := pair{v, c}
			if g[p] == nil {
				q := pair{v, 0}
				g[p] = []pair{q}
				g[q] = append(g[q], p)
			}
		}
		for ; m > 0; m-- {
			var v, w, c int
			Fscan(in, &v, &w, &c)
			p := pair{v, c}
			q := pair{w, c}
			add(v, c)
			add(w, c)
			g[p] = append(g[p], q)
			g[q] = append(g[q], p)
		}
		Fscan(in, &st, &end)

		dis := make(map[pair]int, len(g))
		for p := range g {
			dis[p] = 1e9
		}
		dis[pair{st, 0}] = 0
		type vd struct{ v pair; d int }
		ql, qr := []vd{{pair{st, 0}, 0}}, []vd{}
		for {
			var p vd
			if len(ql) > 0 {
				ql, p = ql[:len(ql)-1], ql[len(ql)-1]
			} else {
				p, qr = qr[0], qr[1:]
			}
			v := p.v
			if v.v == end {
				Fprintln(out, p.d)
				break
			}
			if p.d > dis[v] {
				continue
			}
			for _, w := range g[v] {
				newD := p.d
				if v.c == 0 {
					newD++
				}
				if newD < dis[w] {
					dis[w] = newD
					if v.c > 0 {
						ql = append(ql, vd{w, newD})
					} else {
						qr = append(qr, vd{w, newD})
					}
				}
			}
		}
	}
}

//func main() { cf1941G(bufio.NewReader(os.Stdin), os.Stdout) }

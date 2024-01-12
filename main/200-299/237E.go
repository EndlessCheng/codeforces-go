package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf237E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int = 1e18

	var s string
	var n, lim int
	Fscan(in, &s, &n)
	ns := len(s)

	st := 26 + n
	end := st + 1

	type nb struct{ to, rid, cap, cost int }
	g := make([][]nb, end+1)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap, cost})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0, -cost})
	}
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	for i, c := range cnt {
		addEdge(st, i, c, 0)
	}
	for j := 26; j < 26+n; j++ {
		Fscan(in, &s, &lim)
		cnt = [26]int{}
		for _, b := range s {
			cnt[b-'a']++
		}
		addEdge(j, end, lim, 0)
		for i, c := range cnt {
			addEdge(i, j, c, j-25)
		}
	}

	dist := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dist {
			dist[i] = inf
		}
		dist[st] = 0
		inQ[st] = true
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + e.cost; newD < dist[w] {
					dist[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		return dist[end] < inf
	}
	edmondsKarp := func() (maxFlow, minCost int) {
		for spfa() {
			minF := inf
			for v := end; v != st; {
				p := fa[v]
				if c := g[p.v][p.i].cap; c < minF {
					minF = c
				}
				v = p.v
			}
			for v := end; v != st; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap -= minF
				g[v][e.rid].cap += minF
				v = p.v
			}
			maxFlow += minF
			minCost += dist[end] * minF
		}
		return
	}
	mf, ans := edmondsKarp()
	if mf < ns {
		Fprint(out, -1)
	} else {
		Fprint(out, ans)
	}
}

//func main() { cf237E(os.Stdin, os.Stdout) }

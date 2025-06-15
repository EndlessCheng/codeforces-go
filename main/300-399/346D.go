package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf346D(in io.Reader, out io.Writer) {
	var n, m, s, t int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	deg := make([]int, n+1)
	for range m {
		var v, w int
		Fscan(in, &v, &w)
		g[w] = append(g[w], v)
		deg[v]++
	}
	Fscan(in, &s, &t)

	dis := make([]int, n+1)
	for i := range dis {
		dis[i] = 1e9
	}
	dis[t] = 0
	type vd struct{ v, d int }
	q := [2][]vd{{{t, 0}}}
	for len(q[0]) > 0 || len(q[1]) > 0 {
		var p vd
		if len(q[0]) > 0 {
			q[0], p = q[0][:len(q[0])-1], q[0][len(q[0])-1]
		} else {
			p, q[1] = q[1][0], q[1][1:]
		}
		v := p.v
		if p.d > dis[v] {
			continue
		}
		for _, w := range g[v] {
			wt := 0
			deg[w]--
			if deg[w] > 0 {
				wt = 1
			}
			newD := p.d + wt
			if newD < dis[w] {
				dis[w] = newD
				q[wt] = append(q[wt], vd{w, newD})
			}
		}
	}
	ans := dis[s]
	if ans == 1e9 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { cf346D(bufio.NewReader(os.Stdin), os.Stdout) }

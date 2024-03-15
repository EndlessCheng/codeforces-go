package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf527E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	type edge struct{ to, eid int }
	g := make([][]edge, n)
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], edge{w, i})
		g[w] = append(g[w], edge{v, i})
	}

	pre := -1
	for i, vs := range g {
		if len(vs)%2 == 0 {
			continue
		}
		if pre < 0 {
			pre = i
		} else {
			g[pre] = append(g[pre], edge{i, m})
			g[i] = append(g[i], edge{pre, m})
			m++
			pre = -1
		}
	}
	if m%2 > 0 {
		g[0] = append(g[0], edge{0, m})
		m++
	}

	Fprintln(out, m)
	rev := false
	vis := make([]bool, m)
	var f func(int)
	f = func(v int) {
		for len(g[v]) > 0 {
			e := g[v][0]
			g[v] = g[v][1:]
			i := e.eid
			if vis[i] {
				continue
			}
			vis[i] = true
			w := e.to
			f(w)
			if rev {
				Fprintln(out, w+1, v+1)
			} else {
				Fprintln(out, v+1, w+1)
			}
			rev = !rev
		}
	}
	f(0)
}

//func main() { cf527E(os.Stdin, os.Stdout) }

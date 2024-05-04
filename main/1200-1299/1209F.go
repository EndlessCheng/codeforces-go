package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1209F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	n0 := n
	type nb struct{ to, wt int }
	g := [1e6][]nb{}
	add := func(v, w, x int) {
		for ; x > 9; x /= 10 {
			g[n] = []nb{{w, x % 10}}
			w = n
			n++
		}
		g[v] = append(g[v], nb{w, x})
	}
	for i := 1; i <= m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		add(v, w, i)
		add(w, v, i)
	}

	dis := make([]int, n)
	vis := make([]bool, n)
	vis[0] = true
	q := [][]int{{0}}
	for len(q) > 0 {
		vs := q[0]
		q = q[1:]
		type edge struct{ from, to int }
		nxt := [10][]edge{}
		for _, v := range vs {
			for _, e := range g[v] {
				nxt[e.wt] = append(nxt[e.wt], edge{v, e.to})
			}
		}
		for wt, es := range nxt {
			ws := []int{}
			for _, e := range es {
				w := e.to
				if !vis[w] {
					vis[w] = true
					dis[w] = (dis[e.from]*10 + wt) % 1_000_000_007
					ws = append(ws, w)
				}
			}
			if len(ws) > 0 {
				q = append(q, ws)
			}
		}
	}
	for _, v := range dis[1:n0] {
		Fprintln(out, v)
	}
}

//func main() { cf1209F(os.Stdin, os.Stdout) }

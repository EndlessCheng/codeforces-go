package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf911F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, u int
	Fscan(in, &n)
	g := make([][]int, n)
	deg := make([]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		deg[v]++
		deg[w]++
	}

	maxD := -1
	var dfs func(int, int, int)
	dfs = func(v, fa, d int) {
		if d > maxD {
			maxD, u = d, v
		}
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, d+1)
			}
		}
	}
	dfs(0, -1, 0)
	dv := u
	maxD = -1
	dfs(u, -1, 0)
	dw := u
	ans := maxD * (maxD + 1) / 2

	f := make([]struct{ v, d int }, n)
	for i := range f {
		f[i].d = -1
	}
	var findFarthest func(int, int, int, int)
	findFarthest = func(v, fa, d, tar int) {
		if d > f[v].d {
			f[v].d = d
			f[v].v = tar
		}
		for _, w := range g[v] {
			if w != fa {
				findFarthest(w, v, d+1, tar)
			}
		}
	}
	findFarthest(dv, -1, 0, dv)
	findFarthest(dw, -1, 0, dw)

	op := [][3]int{}
	q := []int{}
	for i, d := range deg {
		if d == 1 && i != dv && i != dw {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		p := &f[v]
		ans += p.d
		p.d = -1
		op = append(op, [3]int{v, p.v, v})
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 1 {
				q = append(q, w)
			}
		}
	}

	for v := dv; v != dw; {
		f[v].d = -1
		op = append(op, [3]int{v, dw, v})
		for _, w := range g[v] {
			if f[w].d >= 0 {
				v = w
				break
			}
		}
	}
	Fprintln(out, ans)
	for _, p := range op {
		Fprintln(out, p[0]+1, p[1]+1, p[2]+1)
	}
}

//func main() { cf911F(os.Stdin, os.Stdout) }

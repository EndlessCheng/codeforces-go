package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1037E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type edge struct{ v, w int }

	var n, m, k, v, w, cnt int
	Fscan(in, &n, &m, &k)
	g := make([][]int, n)
	e := make([]edge, m)
	for i := range e {
		Fscan(in, &v, &w)
		v--
		w--
		e[i] = edge{v, w}
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	d := make([]int, n)
	q := []int{}
	for i, e := range g {
		d[i] = len(e)
		if d[i] < k {
			q = append(q, i)
		} else {
			cnt++
		}
	}

	var del func(int)
	del = func(v int) {
		if d[v]--; d[v] == k-1 {
			cnt--
			for _, w := range g[v] {
				del(w)
			}
		}
	}
	for _, v := range q {
		for _, w := range g[v] {
			del(w)
		}
	}

	ans := make([]int, m)
	for i := m - 1; i >= 0; i-- {
		ans[i] = cnt
		v, w := e[i].v, e[i].w
		if d[v] >= k {
			g[v] = g[v][:len(g[v])-1]
			del(w)
		}
		if d[w] >= k {
			g[w] = g[w][:len(g[w])-1]
			del(v)
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF1037E(os.Stdin, os.Stdout) }

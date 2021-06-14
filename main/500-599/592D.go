package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF592D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	d := make([]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		d[v]++
		d[w]++
	}
	b := make([]bool, n)
	ignore := make([]bool, n)
	for ; m > 0; m-- {
		Fscan(in, &v)
		v--
		b[v] = true
	}

	left := n
	q := []int{}
	for i, b := range b {
		if !b && d[i] == 1 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		ignore[v] = true
		left--
		for _, w := range g[v] {
			if d[w]--; d[w] == 1 && !b[w] {
				q = append(q, w)
			}
		}
	}

	maxD, u := -1, 0
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > maxD {
			maxD, u = d, v
		}
		for _, w := range g[v] {
			if w != fa && !ignore[w] {
				f(w, v, d+1)
			}
		}
	}
	// 这里的 v 是最后一个读入的 v
	f(v, -1, 0)
	dv := u
	maxD = -1
	f(u, -1, 0)
	dw := u

	minEnd := n - 1
	var findEnds func(v, fa, d int)
	findEnds = func(v, fa, d int) {
		if d == maxD {
			if v < minEnd {
				minEnd = v
			}
			return
		}
		for _, w := range g[v] {
			if w != fa && !ignore[w] {
				findEnds(w, v, d+1)
			}
		}
	}
	findEnds(dv, -1, 0)
	findEnds(dw, -1, 0)
	Fprintln(out, minEnd+1)
	Fprintln(out, (left-1)*2-maxD) // 减一是点转边
}

//func main() { CF592D(os.Stdin, os.Stdout) }

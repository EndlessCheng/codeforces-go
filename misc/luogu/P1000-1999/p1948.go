package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func p1948(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, v, w, wt int
	Fscan(in, &n, &m, &k)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}
	var l, r []int
	pop := func() (v int) {
		if len(l) > 0 {
			l, v = l[:len(l)-1], l[len(l)-1]
		} else {
			v, r = r[0], r[1:]
		}
		return
	}

	const inf int = 1e9
	dist := make([]int, n)
	ans := sort.Search(1e6+1, func(c int) bool {
		for i := range dist {
			dist[i] = inf
		}
		dist[0] = 0
		l, r = []int{0}, nil
		for len(l) > 0 || len(r) > 0 {
			v := pop()
			for _, e := range g[v] {
				w := e.to
				d := 0
				if e.wt > c {
					d = 1
				}
				if newD := dist[v] + d; newD < dist[w] {
					dist[w] = newD
					if d == 0 {
						l = append(l, w)
					} else {
						r = append(r, w)
					}
				}
			}
		}
		return dist[n-1] <= k
	})
	if ans > 1e6 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { p1948(os.Stdin, os.Stdout) }

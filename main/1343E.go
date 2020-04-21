package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1343E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, m, a, b, c, v, w int
	calcDep := func(g [][]int, st int) []int {
		d := make([]int, n+1)
		for i := range d {
			d[i] = -1
		}
		d[st] = 0
		for q := []int{st}; len(q) > 0; {
			v, q = q[0], q[1:]
			for _, w := range g[v] {
				if d[w] < 0 {
					d[w] = d[v] + 1
					q = append(q, w)
				}
			}
		}
		return d
	}
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m, &a, &b, &c)
		p := make([]int, m)
		for i := range p {
			Fscan(in, &p[i])
		}
		sort.Ints(p)
		sum := make([]int64, m+1)
		for i, v := range p {
			sum[i+1] = sum[i] + int64(v)
		}
		g := make([][]int, n+1)
		for i := 0; i < m; i++ {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		ans := int64(1e18)
		depA, depB, depC := calcDep(g, a), calcDep(g, b), calcDep(g, c)
		for v := 1; v <= n; v++ {
			if cost := depA[v] + depB[v] + depC[v]; cost <= m {
				if s := sum[cost] + sum[depB[v]]; s < ans {
					ans = s
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1343E(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF796D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, d, v, w int
	Fscan(in, &n, &k, &d)
	type pair struct{ v, fa int }
	q := []pair{}
	vis := make([]bool, n+1)
	for ; k > 0; k-- {
		Fscan(in, &v)
		if !vis[v] {
			vis[v] = true
			q = append(q, pair{v, 0})
		}
	}
	type nb struct{ to, i int }
	g := make([][]nb, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], nb{w, i})
		g[w] = append(g[w], nb{v, i})
	}

	ans := []interface{}{}
	inAns := make([]bool, n)
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, e := range g[p.v] {
			if w := e.to; w != p.fa {
				if vis[w] {
					if !inAns[e.i] {
						ans = append(ans, e.i)
						inAns[e.i] = true
					}
				} else {
					vis[w] = true
					q = append(q, pair{w, p.v})
				}
			}
		}
	}
	Fprintln(out, len(ans))
	Fprintln(out, ans...)
}

//func main() { CF796D(os.Stdin, os.Stdout) }

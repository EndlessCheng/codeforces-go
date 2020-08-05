package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF263D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k, v, w, end, st int
	Fscan(in, &n, &m, &k)
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	fa := make([]int, n+1)
	dep := make([]int, n+1)
	var f func(v, p, d int) bool
	f = func(v, p, d int) bool {
		fa[v] = p
		dep[v] = d
		for _, w := range g[v] {
			if dep[w] == 0 {
				if f(w, v, d+1) {
					return true
				}
			} else if d-dep[w] >= k {
				end, st = v, w
				return true
			}
		}
		return false
	}
	f(1, 0, 1)

	ans := []interface{}{st}
	for v := end; v != st; v = fa[v] {
		ans = append(ans, v)
	}
	Fprintln(out, len(ans))
	Fprint(out, ans...)
}

//func main() { CF263D(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1627C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type edge struct{ to, i int }
		g := make([][]edge, n)
		ok := true
		for i := 0; i < n-1; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], edge{w, i})
			g[w] = append(g[w], edge{v, i})
			ok = ok && len(g[v]) < 3 && len(g[w]) < 3
		}
		if !ok {
			Fprintln(out, -1)
			continue
		}
		st := 0
		for i, e := range g {
			if len(e) == 1 {
				st = i
				break
			}
		}
		ans := make([]int, n-1)
		cur := 2
		var f func(int, int)
		f = func(v, fa int) {
			for _, e := range g[v] {
				w := e.to
				if w != fa {
					ans[e.i] = cur
					cur = 5 - cur
					f(w, v)
				}
			}
		}
		f(st, -1)
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1627C(os.Stdin, os.Stdout) }

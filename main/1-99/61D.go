package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF61D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var v, w, wt, ans, mx uint
	Fscan(in, &n)
	type edge struct{ to, wt uint }
	g := make([][]edge, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &wt)
		g[v] = append(g[v], edge{w, wt})
		g[w] = append(g[w], edge{v, wt})
		ans += 2 * wt
	}
	var f func(v, fa, s uint)
	f = func(v, fa, s uint) {
		if s > mx {
			mx = s
		}
		for _, e := range g[v] {
			if w := e.to; w != fa {
				f(w, v, s+e.wt)
			}
		}
	}
	f(1, 0, 0)
	Fprint(out, ans-mx)
}

//func main() { CF61D(os.Stdin, os.Stdout) }
